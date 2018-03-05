package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/matrix/globalTypes"
	"github.com/Nordgedanken/Morpheus/matrix/messages"
	"github.com/Nordgedanken/Morpheus/matrix/rooms"
	"github.com/Nordgedanken/Morpheus/matrix/syncer"
	"github.com/Nordgedanken/Morpheus/ui/listLayouts"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/opennota/linkify"
	"github.com/pkg/errors"
	"github.com/rhinoman/go-commonmark"
	"github.com/shibukawa/configdir"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// NewMainUIStruct gives you a MainUI struct with prefilled data
func NewMainUIStruct(windowWidth, windowHeight int, window *widgets.QMainWindow) (mainUIStruct *MainUI) {
	configStruct := globalTypes.Config{
		WindowWidth:  windowWidth,
		WindowHeight: windowHeight,
		Rooms:        make(map[string]*rooms.Room),
	}
	mainUIStruct = &MainUI{
		Config: configStruct,
		window: window,
	}
	return
}

// NewMainUIStructWithExistingConfig gives you a MainUI struct with prefilled data and data from a previous Config
func NewMainUIStructWithExistingConfig(configStruct globalTypes.Config, window *widgets.QMainWindow) (mainUIStruct *MainUI) {
	configStruct.Rooms = make(map[string]*rooms.Room)
	mainUIStruct = &MainUI{
		Config: configStruct,
		window: window,
	}
	return
}

// SetCli allows you to add a gomatrix.Client to your MainUI struct
func (m *MainUI) SetCli(cli *gomatrix.Client) {
	m.Cli = cli
}

// GetWidget gives you the widget of the MainUI struct
func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	widget = m.widget
	return
}

// NewUI initializes a new Main Screen
func (m *MainUI) NewUI() (err error) {
	m.loadChatUIDefaults()

	//Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(m.widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		LogoutErr := m.logout()
		if LogoutErr != nil {
			err = LogoutErr
			return
		}
	})

	m.initScrolls()

	m.MessageList.ConnectTriggerMessage(func(messageID string) {
		log.Infoln("Trigger Message")
		var own bool
		var message = m.Rooms[m.CurrentRoom].Messages[messageID]
		if message.Author == m.Cli.UserID {
			own = true
		} else {
			own = false
		}
		height := m.App.FontMetrics().Height()
		width := m.App.FontMetrics().Width(message.Message, len(message.Message))

		m.MessageList.NewMessage(message, m.messageScrollArea, own, height, width)
	})

	go m.startSync()
	m.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	m.MainWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	m.RoomList.ConnectTriggerRoom(func(roomID string) {
		room := m.Rooms[roomID]

		NewRoomErr := m.RoomList.NewRoom(room, m.roomScrollArea)
		if NewRoomErr != nil {
			err = NewRoomErr
			return
		}
	})

	go m.initRoomList()

	var message string
	messageInput := widgets.NewQLineEditFromPointer(m.widget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
	messageInput.ConnectTextChanged(func(value string) {
		message = value
	})

	m.window.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			go m.sendMessage(message)

			messageInput.Clear()
			ev.Accept()
		} else {
			messageInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
		return
	})

	m.RoomList.ConnectChangeRoom(func(roomID string) {
		room := m.Rooms[roomID]

		if m.CurrentRoom != room.RoomID {
			m.SetCurrentRoom(room.RoomID)
			m.MainWidget.SetWindowTitle("Morpheus - " + room.GetRoomTopic())

			room.ConnectSetAvatar(func(IMGdata []byte) {
				avatar := gui.NewQPixmap()

				str := string(IMGdata[:])
				avatar.LoadFromData(str, uint(len(str)), "", 0)
				m.RoomAvatar.SetPixmap(avatar)

				return
			})

			go room.GetRoomAvatar()

			m.RoomTitle.SetText(room.GetRoomName())

			m.RoomTopic.SetText(room.GetRoomTopic())
			count := m.MessageList.Count()
			for i := 0; i < count; i++ {
				if (i % 10) == 0 {
					m.App.ProcessEvents(core.QEventLoop__AllEvents)
				}
				widgetScroll := m.MessageList.ItemAt(i).Widget()
				widgetScroll.DeleteLater()
			}

			log.Println("Before loadCache")
			go m.loadCache()
		}
	})

	return
}

func (m *MainUI) initScrolls() {
	m.roomScrollArea.SetWidgetResizable(true)
	m.roomScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	m.roomScrollArea.SetContentsMargins(0, 0, 0, 0)
	m.roomScrollArea.SetSizeAdjustPolicy(widgets.QAbstractScrollArea__AdjustToContents)

	m.messageScrollArea.SetWidgetResizable(true)
	m.messageScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	m.messageScrollArea.SetContentsMargins(0, 0, 0, 0)
}

func (m *MainUI) loadChatUIDefaults() {
	m.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	m.MainWidget = loader.Load(file, m.widget)
	file.Close()

	m.messageScrollArea = widgets.NewQScrollAreaFromPointer(m.widget.FindChild("messageScroll", core.Qt__FindChildrenRecursively).Pointer())
	m.roomScrollArea = widgets.NewQScrollAreaFromPointer(m.widget.FindChild("roomScroll", core.Qt__FindChildrenRecursively).Pointer())

	m.MessageList = listLayouts.NewMessageList2(m.messageScrollArea)
	m.MessageList.Init(m.messageScrollArea)
	m.RoomList = listLayouts.NewRoomList2(m.roomScrollArea)
	m.RoomList.Init(m.roomScrollArea)

	m.RoomAvatar = widgets.NewQLabelFromPointer(m.widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	m.RoomTitle = widgets.NewQLabelFromPointer(m.widget.FindChild("RoomTitle", core.Qt__FindChildrenRecursively).Pointer())
	m.RoomTopic = widgets.NewQLabelFromPointer(m.widget.FindChild("Topic", core.Qt__FindChildrenRecursively).Pointer())

	var layout = widgets.NewQHBoxLayout()
	m.window.SetLayout(layout)
	layout.InsertWidget(0, m.MainWidget, 0, core.Qt__AlignTop|core.Qt__AlignLeft)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)

	m.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		m.MainWidget.Resize(event.Size())
		event.Accept()
	})

	//Set Avatar
	/*avatarLogo := widgets.NewQLabelFromPointer(m.widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatar, AvatarErr := matrix.GetOwnUserAvatar(m.Cli)
	if AvatarErr != nil {
		err = AvatarErr
		return
	}
	avatarLogo.SetPixmap(avatar)*/
}

func (m *MainUI) sendMessage(message string) (err error) {
	messageOriginal := message
	lm := linkify.Links(message)
	for _, l := range lm {
		link := message[l.Start:l.End]
		message = strings.Replace(message, link, "<a href='"+link+"'>"+link+"</a>", -1)
	}

	mardownMessage := commonmark.Md2Html(message, 0)
	if mardownMessage == message {
		_, SendErr := m.Cli.SendMessageEvent(m.CurrentRoom, "m.room.message", matrix.HTMLMessage{MsgType: "m.text", Body: messageOriginal, FormattedBody: message, Format: "org.matrix.custom.html"})
		if SendErr != nil {
			err = SendErr
			return
		}
	} else {
		_, SendErr := m.Cli.SendMessageEvent(m.CurrentRoom, "m.room.message", matrix.HTMLMessage{MsgType: "m.text", Body: message, FormattedBody: mardownMessage, Format: "org.matrix.custom.html"})
		if SendErr != nil {
			err = SendErr
			return
		}
	}
	log.Infoln("SendErr: ", err)
	return
}

func (m *MainUI) logout() (err error) {
	log.Infoln("Starting Logout Sequence in background")
	var wg sync.WaitGroup
	results := make(chan bool)

	wg.Add(1)
	go func(cli *gomatrix.Client, results chan<- bool) {
		defer wg.Done()
		cli.StopSync()
		_, LogoutErr := cli.Logout()
		if LogoutErr != nil {
			log.Errorln(LogoutErr)
			results <- false
		}
		cli.ClearCredentials()

		userDB, DBOpenErr := db.OpenUserDB()
		if DBOpenErr != nil {
			log.Errorln(DBOpenErr)
		}
		userDB.Close()

		configDirs := configdir.New("Nordgedanken", "Morpheus")
		filePath := filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)
		DeleteErr := os.RemoveAll(filePath + "/data/user/")
		if DeleteErr != nil {
			log.Errorln(DeleteErr)
			results <- false
		}
		db.ResetOnceUser()
		results <- true
	}(m.Cli, results)

	go func() {
		wg.Wait()      // wait for each execTask to return
		close(results) // then close the results channel
	}()

	//Show LoginUI
	for result := range results {
		if result {
			m.window.DisconnectKeyPressEvent()
			m.window.DisconnectResizeEvent()
			m.widget.DisconnectResizeEvent()
			m.messageScrollArea.DisconnectResizeEvent()

			LoginUIStruct := NewLoginUIStructWithExistingConfig(m.Config, m.window)
			loginUIErr := LoginUIStruct.NewUI()
			if loginUIErr != nil {
				err = loginUIErr
				return
			}
			m.window.SetCentralWidget(LoginUIStruct.GetWidget())
		}
	}
	return
}

func (m *MainUI) startSync() (err error) {
	//Start Syncer!
	m.storage = syncer.NewMorpheusStore()

	Syncer := syncer.NewMorpheusSyncer(m.Cli.UserID, m.storage, &m.Config)

	m.Cli.Store = m.storage
	m.Cli.Syncer = Syncer
	Syncer.Store = m.storage

	Syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		log.Infoln("NewEVENT")
		formattedBody, _ := ev.Content["formatted_body"]
		var msg string
		msg, _ = formattedBody.(string)
		if msg == "" {
			msg, _ = ev.Body()
		}
		room := ev.RoomID
		sender := ev.Sender
		id := ev.ID
		timestamp := ev.Timestamp
		go db.CacheMessageEvents(id, sender, room, msg, timestamp)
		if room == m.CurrentRoom {
			message := messages.NewMessage()
			message.EventID = id
			message.Author = sender
			message.Message = msg
			message.Timestamp = timestamp
			message.Cli = m.Cli
			m.Rooms[room].AddMessage(message)

			go m.MessageList.TriggerMessage(id)
			m.MessageList.MessageCount++

			if (m.MessageList.MessageCount % 10) == 0 {
				m.App.ProcessEvents(core.QEventLoop__AllEvents)
			}
		}
		m.App.ProcessEvents(core.QEventLoop__AllEvents)
	})

	Syncer.OnEventType("m.room.name", func(ev *gomatrix.Event) {
		roomNameRaw, _ := ev.Content["name"]
		var roomName string
		roomName, _ = roomNameRaw.(string)
		evType := ev.Type
		room := ev.RoomID
		go m.Rooms[room].UpdateRoomNameByEvent(roomName, evType)
	})

	// Start Non-blocking sync
	go func() {
		log.Infoln("Start sync")
		for {
			e := m.Cli.Sync()
			if e == nil {
				break
			}
			if e != nil {
				err = e
			}
		}
	}()
	return
}

func (m *MainUI) initRoomList() (err error) {
	roomsStruct, roomsErr := rooms.GetRooms(m.Cli)
	if roomsErr != nil {
		err = roomsErr
		return
	}

	first := true
	for _, roomID := range roomsStruct {
		m.Rooms[roomID] = rooms.NewRoom()
		m.Rooms[roomID].Cli = m.Cli
		m.Rooms[roomID].RoomID = roomID
		go m.RoomList.TriggerRoom(roomID)
		m.RoomList.RoomCount++
		if (m.RoomList.RoomCount % 10) == 0 {
			m.App.ProcessEvents(core.QEventLoop__AllEvents)
		}
		if first {
			m.RoomList.ChangeRoom(roomID)
			m.App.ProcessEvents(core.QEventLoop__AllEvents)
		}
		first = false
	}

	m.App.ProcessEvents(core.QEventLoop__AllEvents)
	return
}

func (m *MainUI) loadCache() (err error) {
	log.Println("Loading cache!")
	barAtBottom := false
	bar := m.messageScrollArea.VerticalScrollBar()
	if bar.Value() == bar.Maximum() {
		barAtBottom = true
	}

	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
	}
	log.Infoln("room|" + m.CurrentRoom + "|messages|")
	MsgPrefix := []byte("room|" + m.CurrentRoom + "|messages|")
	DBerr := cacheDB.View(func(txn *badger.Txn) error {
		log.Println("CacheDB")
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		MsgIt := txn.NewIterator(opts)

		//DEBUG
		debugResult, QueryErr := db.Get(txn, []byte("room|"+m.CurrentRoom+"|messages|"))
		if QueryErr != nil {
			return errors.WithMessage(QueryErr, "Key: "+"room|"+m.CurrentRoom+"|messages|")
		}
		log.Infoln("Debug Result: ", debugResult)

		doneMsg := make(map[string]bool)
		valid := func() bool {
			log.Infoln("Item: ", MsgIt.Item())
			valid := MsgIt.ValidForPrefix(MsgPrefix)
			log.Println("Valid: ", valid)
			return valid
		}()

		for MsgIt.Seek(MsgPrefix); valid; MsgIt.Next() {
			log.Println("MSG LOOP")
			item := MsgIt.Item()
			key := item.Key()
			stringKey := fmt.Sprintf("%s", key)
			stringKeySlice := strings.Split(stringKey, "|")
			stringKeyEnd := stringKeySlice[len(stringKeySlice)-1]
			if stringKeyEnd != "id" {
				continue
			}

			value, ValueErr := item.Value()
			if ValueErr != nil {
				return ValueErr
			}
			idValue := fmt.Sprintf("%s", value)

			if !doneMsg[idValue] {
				// Remember we already added this message to the view
				doneMsg[idValue] = true

				// Get all Data
				senderResult, QueryErr := db.Get(txn, []byte(strings.Replace(stringKey, "|id", "|sender", -1)))
				if QueryErr != nil {
					return errors.WithMessage(QueryErr, "Key: "+strings.Replace(stringKey, "|id", "|sender", -1))
				}
				sender := fmt.Sprintf("%s", senderResult)

				msgResult, QueryErr := db.Get(txn, []byte(strings.Replace(stringKey, "|id", "|messageString", -1)))
				if QueryErr != nil {
					return errors.WithMessage(QueryErr, "Key: "+strings.Replace(stringKey, "|id", "|messageString", -1))
				}
				msg := fmt.Sprintf("%s", msgResult)

				timestampResult, QueryErr := db.Get(txn, []byte(strings.Replace(stringKey, "|id", "|timestamp", -1)))
				if QueryErr != nil {
					return errors.WithMessage(QueryErr, "Key: "+strings.Replace(stringKey, "|id", "|timestamp", -1))
				}
				timestamp := fmt.Sprintf("%s", timestampResult)

				timestampInt, ConvErr := strconv.ParseInt(timestamp, 10, 64)
				if ConvErr != nil {
					return errors.WithMessage(ConvErr, "Timestamp String: "+timestamp)
				}

				//TODO Use for better/faster cache loading
				currentRoomMem := m.Rooms[m.CurrentRoom]

				message := messages.NewMessage()
				message.EventID = idValue
				message.Author = sender
				message.Message = msg
				message.Timestamp = timestampInt
				message.Cli = m.Cli
				currentRoomMem.AddMessage(message)

				go m.MessageList.TriggerMessage(idValue)
				m.MessageList.MessageCount++

				log.Println(m.MessageList.MessageCount)

				if (m.MessageList.MessageCount % 10) == 0 {
					m.App.ProcessEvents(core.QEventLoop__AllEvents)
				}
			}
		}
		m.App.ProcessEvents(core.QEventLoop__AllEvents)
		return nil
	})
	if DBerr != nil {
		log.Errorln("DBERR: ", DBerr)
		err = DBerr
		return
	}

	if barAtBottom {
		bar.Update()
		bar.SetValue(bar.Maximum())
	}

	return
}
