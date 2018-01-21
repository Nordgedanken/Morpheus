package ui

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/matrix/globalTypes"
	"github.com/Nordgedanken/Morpheus/matrix/rooms"
	"github.com/Nordgedanken/Morpheus/matrix/syncer"
	"github.com/Nordgedanken/Morpheus/ui/listLayouts"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/opennota/linkify"
	"github.com/pkg/errors"
	"github.com/rhinoman/go-commonmark"
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

	//Set Avatar
	avatarLogo := widgets.NewQLabelFromPointer(m.widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatar, AvatarErr := matrix.GetOwnUserAvatar(m.Cli)
	if AvatarErr != nil {
		err = AvatarErr
		return
	}
	avatarLogo.SetPixmap(avatar)

	//Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(m.widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		LogoutErr := m.logout(m.widget, m.messageScrollArea)
		if LogoutErr != nil {
			err = LogoutErr
			return
		}
	})

	m.initScrolls()

	m.MessageListLayout.ConnectTriggerMessage(func(messageBody, sender string, timestamp int64) {
		log.Println("triggered Message")
		var own bool
		if sender == m.Cli.UserID {
			own = true
		} else {
			own = false
		}
		lm := linkify.Links(messageBody)
		for _, l := range lm {
			link := messageBody[l.Start:l.End]
			if l.Start-9 > 0 {
				if !strings.Contains(messageBody[l.Start-9:l.Start], "<a href='") {
					if l.Start-(1+l.Start+l.End) > 0 {
						if !strings.Contains(messageBody[l.Start-(1+l.Start+l.End):l.Start], "<a href='"+link+"'>") {
							messageBody = strings.Replace(messageBody, link, "<a href='"+link+"'>"+link+"</a>", -1)
						}
					} else if l.Start-(1+l.Start+l.End) <= 0 {
						if !strings.Contains(messageBody[0:l.Start], "<a href='"+link+"'>") {
							messageBody = strings.Replace(messageBody, link, "<a href='"+link+"'>"+link+"</a>", -1)
						}
					}
				}
			} else if l.Start-9 <= 0 {
				if !strings.Contains(messageBody[0:l.Start], "<a href='") {
					if !strings.Contains(messageBody[0:l.Start], "<a href='"+link+"'>") {
						messageBody = strings.Replace(messageBody, link, "<a href='"+link+"'>"+link+"</a>", -1)
					}
				}
			}
		}
		m.MessageListLayout.NewMessage(messageBody, m.Cli, sender, timestamp, m.messageScrollArea, own)
	})

	go m.startSync()
	m.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	m.MainWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	m.RoomListLayout.ConnectTriggerRoom(func(roomID string) {
		room := m.Rooms[roomID]

		NewRoomErr := m.RoomListLayout.NewRoom(room, m.roomScrollArea)
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

	m.RoomListLayout.ConnectChangeRoom(func(roomID string) {
		room := m.Rooms[roomID]
		roomAvatar, roomAvatarErr := room.GetRoomAvatar()
		if roomAvatarErr != nil {
			err = roomAvatarErr
			return
		}
		if m.CurrentRoom != room.RoomID {
			m.SetCurrentRoom(room.RoomID)
			m.MainWidget.SetWindowTitle("Morpheus - " + room.GetRoomTopic())

			m.RoomAvatar.SetPixmap(roomAvatar)

			m.RoomTitle.SetText(room.GetRoomName())

			m.RoomTopic.SetText(room.GetRoomTopic())
			count := m.MessageListLayout.Count()
			for i := 0; i < count; i++ {
				widgetScroll := m.MessageListLayout.ItemAt(i).Widget()
				widgetScroll.DeleteLater()
			}

			log.Println("next loadCache")
			go m.loadCache()
		}
	})

	return
}

func (m *MainUI) initScrolls() {
	// Init Message View
	m.MessageListLayout = listLayouts.NewMessageList(m.messageScrollArea)

	// Init Room View
	m.RoomListLayout = listLayouts.NewRoomList(m.roomScrollArea)

	m.messageScrollArea.SetWidgetResizable(true)
	m.messageScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	m.messageScrollArea.SetContentsMargins(0, 0, 0, 0)

	m.roomScrollArea.SetWidgetResizable(true)
	m.roomScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	m.roomScrollArea.SetContentsMargins(0, 0, 0, 0)
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
	return
}

func (m *MainUI) logout(widget *widgets.QWidget, messageScrollArea *widgets.QScrollArea) (err error) {
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

		//Flush complete DB
		txn := userDB.NewTransaction(true) // Read-write txn
		QueryErr := txn.Delete([]byte(""))
		if QueryErr != nil {
			log.Errorln(QueryErr)
			results <- false
		}

		CommitErr := txn.Commit(nil)
		if CommitErr != nil {
			log.Errorln(CommitErr)
			results <- false
		}

		DBPurgeErr := userDB.PurgeOlderVersions()
		if DBPurgeErr != nil {
			log.Errorln(DBPurgeErr)
			results <- false
		} else {
			results <- true
		}
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
			widget.DisconnectResizeEvent()
			messageScrollArea.DisconnectResizeEvent()

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
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
	}

	m.SetCacheDB(&db.MorpheusStorage{
		Database: CacheDB,
	})
	m.storage = &syncer.MorpheusStore{
		InMemoryStore: *gomatrix.NewInMemoryStore(),
		CacheDatabase: m.CacheDB,
	}

	Syncer := syncer.NewMorpheusSyncer(m.Cli.UserID, m.storage)

	m.Cli.Store = m.storage
	m.Cli.Syncer = Syncer
	Syncer.Store = m.storage

	Syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
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
			go m.MessageListLayout.TriggerMessage(msg, sender, timestamp)
		}
	})

	Syncer.OnEventType("m.room.name", func(ev *gomatrix.Event) {
		roomNameRaw, _ := ev.Content["name"]
		var roomName string
		roomName, _ = roomNameRaw.(string)
		evType := ev.Type
		room := ev.RoomID
		go m.Rooms[room].UpdateRoomNameByEvent(roomName, evType)
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
		m.Rooms[roomID] = rooms.NewRoom(roomID, m.Cli)
		m.RoomListLayout.TriggerRoom(roomID)
		if first {
			go m.RoomListLayout.ChangeRoom(roomID)
		}
		first = false
	}

	return
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (m *MainUI) loadCache() (err error) {
	barAtBottom := false
	bar := m.messageScrollArea.VerticalScrollBar()
	if bar.Value() == bar.Maximum() {
		barAtBottom = true
	}

	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
	}

	DBerr := cacheDB.View(func(txn *badger.Txn) error {
		MsgOpts := badger.DefaultIteratorOptions
		MsgOpts.PrefetchSize = 10
		MsgIt := txn.NewIterator(MsgOpts)
		MsgPrefix := []byte("room|" + m.CurrentRoom + "|messages|id")

		var doneMsg []string

		for MsgIt.Seek(MsgPrefix); MsgIt.ValidForPrefix(MsgPrefix); MsgIt.Next() {
			item := MsgIt.Item()
			key := item.Key()
			stringKey := fmt.Sprintf("%s", key)

			value, ValueErr := item.Value()
			if ValueErr != nil {
				return ValueErr
			}
			idValue := fmt.Sprintf("%s", value)

			if !contains(doneMsg, idValue) {
				// Remember we already added this message to the view
				doneMsg = append(doneMsg, idValue)

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

				log.Println("next TriggerMessage")
				go m.MessageListLayout.TriggerMessage(msg, sender, timestampInt)
			}
		}

		return nil
	})
	if DBerr != nil {
		log.Errorln("DBERR: ", DBerr)
		err = DBerr
		return
	}

	if barAtBottom {
		bar.SetValue(bar.Maximum())
	}

	return
}
