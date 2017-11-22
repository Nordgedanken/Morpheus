package ui

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/util"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/pkg/errors"
	"github.com/rhinoman/go-commonmark"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// NewMainUIStruct gives you a MainUI struct with prefilled data
func NewMainUIStruct(windowWidth, windowHeight int, window *widgets.QMainWindow) (mainUIStruct MainUI) {
	configStruct := config{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
	}
	mainUIStruct = MainUI{
		config: configStruct,
		window: window,
		rooms:  make(map[string]*matrix.Room),
	}
	return
}

// NewMainUIStructWithExistingConfig gives you a MainUI struct with prefilled data and data from a previous Config
func NewMainUIStructWithExistingConfig(configStruct config, window *widgets.QMainWindow) (mainUIStruct MainUI) {
	mainUIStruct = MainUI{
		config: configStruct,
		window: window,
		rooms:  make(map[string]*matrix.Room),
	}
	return
}

// SetCli allows you to add a gomatrix.Client to your MainUI struct
func (m *MainUI) SetCli(cli *gomatrix.Client) {
	m.cli = cli
}

// GetWidget gives you the widget of the MainUI struct
func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	widget = m.widget
	return
}

// InitLogger adds a new logger to the MainUI struct
func (m *MainUI) InitLogger() error {
	localLog := util.Logger()
	localLog, _, err := util.StartFileLog(localLog)
	if err != nil {
		return err
	}
	m.localLog = localLog
	return nil
}

// NewUI initializes a new Main Screen
func (m *MainUI) NewUI() (err error) {
	m.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	m.MainWidget = loader.Load(file, m.widget)
	file.Close()

	m.messageScrollArea = widgets.NewQScrollAreaFromPointer(m.widget.FindChild("messageScroll", core.Qt__FindChildrenRecursively).Pointer())
	messagesScrollAreaContent := widgets.NewQWidgetFromPointer(m.widget.FindChild("messagesScrollAreaContent", core.Qt__FindChildrenRecursively).Pointer())
	roomScrollArea := widgets.NewQScrollAreaFromPointer(m.widget.FindChild("roomScroll", core.Qt__FindChildrenRecursively).Pointer())
	roomScrollAreaContent := widgets.NewQWidgetFromPointer(m.widget.FindChild("roomScrollAreaContent", core.Qt__FindChildrenRecursively).Pointer())

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
	avatarLogo := widgets.NewQLabelFromPointer(m.widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatar, AvatarErr := matrix.GetOwnUserAvatar(m.cli)
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

	// Init Message View
	m.MessageListLayout = NewMessageList(m.messageScrollArea, messagesScrollAreaContent)

	// Init Room View
	roomListLayout := NewRoomList(roomScrollArea, roomScrollAreaContent)

	m.messageScrollArea.SetWidgetResizable(true)
	m.messageScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	m.messageScrollArea.SetContentsMargins(0, 0, 0, 0)
	//messageScrollArea.SetSizeAdjustPolicy(widgets.QAbstractScrollArea__AdjustToContents)

	roomScrollArea.SetWidgetResizable(true)
	roomScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	roomScrollArea.SetContentsMargins(0, 0, 0, 0)
	roomScrollArea.SetSizeAdjustPolicy(widgets.QAbstractScrollArea__AdjustToContents)

	m.MessageListLayout.ConnectTriggerMessage(func(messageBody, sender string, timestamp int64) {
		var own bool
		if sender == m.cli.UserID {
			own = true
		} else {
			own = false
		}
		NewMessageErr := m.MessageListLayout.NewMessage(messageBody, m.cli, sender, timestamp, m.messageScrollArea, own)
		if NewMessageErr != nil {
			err = NewMessageErr
			return
		}
	})

	m.startSync()
	m.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	m.MainWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	roomListLayout.ConnectTriggerRoom(func(roomID string) {
		room := m.rooms[roomID]

		NewRoomErr := roomListLayout.NewRoom(room, roomScrollArea, m)
		if NewRoomErr != nil {
			err = NewRoomErr
			return
		}
	})

	m.initRoomList(roomListLayout, roomScrollArea)

	go m.loadCache()

	m.MainWidget.SetWindowTitle("Morpheus - " + m.rooms[m.currentRoom].GetRoomTopic())

	avatar, roomAvatarErr := m.rooms[m.currentRoom].GetRoomAvatar()
	if roomAvatarErr != nil {
		err = roomAvatarErr
		return
	}
	m.RoomAvatar.SetPixmap(avatar)

	m.RoomTitle.SetText(m.rooms[m.currentRoom].GetRoomName())

	m.RoomTopic.SetText(m.rooms[m.currentRoom].GetRoomTopic())

	var message string
	messageInput := widgets.NewQLineEditFromPointer(m.widget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
	messageInput.ConnectTextChanged(func(value string) {
		message = value
	})

	m.window.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			MessageErr := m.sendMessage(message)
			if MessageErr != nil {
				err = MessageErr
				return
			}

			messageInput.Clear()
			ev.Accept()
		} else {
			messageInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})

	return
}

func (m *MainUI) sendMessage(message string) (err error) {
	mardownMessage := commonmark.Md2Html(message, 0)
	if mardownMessage == message {
		_, SendErr := m.cli.SendText(m.currentRoom, message)
		if SendErr != nil {
			err = SendErr
			return
		}
	} else {
		_, SendErr := m.cli.SendMessageEvent(m.currentRoom, "m.room.message", matrix.HTMLMessage{MsgType: "m.text", Body: message, FormattedBody: mardownMessage, Format: "org.matrix.custom.html"})
		if SendErr != nil {
			err = SendErr
			return
		}
	}
	return
}

func (m *MainUI) logout(widget *widgets.QWidget, messageScrollArea *widgets.QScrollArea) (err error) {
	//TODO register enter and show loader or so
	m.localLog.Println("Starting Logout Sequence in background")
	var wg sync.WaitGroup
	results := make(chan bool)

	wg.Add(1)
	go func(cli *gomatrix.Client, localLog *log.Logger, results chan<- bool) {
		defer wg.Done()
		_, LogoutErr := cli.Logout()
		if LogoutErr != nil {
			localLog.Println(LogoutErr)
			results <- false
		}
		cli.ClearCredentials()

		db, DBOpenErr := db.OpenUserDB()
		if DBOpenErr != nil {
			localLog.Fatalln(DBOpenErr)
		}

		//Flush complete DB
		txn := db.NewTransaction(true) // Read-write txn
		QueryErr := txn.Delete([]byte(""))
		if QueryErr != nil {
			localLog.Println(QueryErr)
			results <- false
		}

		CommitErr := txn.Commit(nil)
		if CommitErr != nil {
			localLog.Println(CommitErr)
			results <- false
		}

		DBPurgeErr := db.PurgeOlderVersions()
		if DBPurgeErr != nil {
			localLog.Println(DBPurgeErr)
			results <- false
		} else {
			results <- true
		}
	}(m.cli, m.localLog, results)

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

			LoginUIStruct := NewLoginUIStructWithExistingConfig(m.config, m.window)
			LoginUIStructInitErr := LoginUIStruct.InitLogger()
			if LoginUIStructInitErr != nil {
				err = LoginUIStructInitErr
				return
			}
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
	m.syncer = m.cli.Syncer.(*gomatrix.DefaultSyncer)
	m.storage = gomatrix.NewInMemoryStore()
	m.cli.Store = m.storage
	m.syncer.Store = m.storage

	m.syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
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
		if room == m.currentRoom {
			go m.MessageListLayout.TriggerMessage(msg, sender, timestamp)
		}
	})

	// Start Non-blocking sync
	go func() {
		m.localLog.Println("Start sync")
		for {

			if e := m.cli.Sync(); e != nil {
				err = e
			}
		}
	}()
	return
}

func (m *MainUI) initRoomList(roomListLayout *QRoomVBoxLayoutWithTriggerSlot, roomScrollArea *widgets.QScrollArea) (err error) {
	rooms, ReqErr := m.cli.JoinedRooms()
	if ReqErr != nil {
		err = ReqErr
		return
	}

	x := 0
	for _, roomID := range rooms.JoinedRooms {
		if x == 0 {
			m.currentRoom = roomID
		}
		x++
		m.rooms[roomID] = matrix.NewRoom(roomID, m.cli)
		roomListLayout.TriggerRoom(roomID)
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
	/*barAtBottom := false
	bar := m.messageScrollArea.VerticalScrollBar()
	if bar.Value() == bar.Maximum() {
		barAtBottom = true
	}*/

	db, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
	}

	log.Println("load cache")

	DBerr := db.View(func(txn *badger.Txn) error {
		MsgOpts := badger.DefaultIteratorOptions
		MsgOpts.PrefetchSize = 10
		MsgIt := txn.NewIterator(MsgOpts)
		MsgPrefix := []byte("room|" + m.currentRoom + "|messages|")

		var doneMsg []string

		for MsgIt.Seek(MsgPrefix); MsgIt.ValidForPrefix(MsgPrefix); MsgIt.Next() {
			item := MsgIt.Item()
			key := item.Key()
			stringKey := fmt.Sprintf("%s", key)
			value, ValueErr := item.Value()
			if ValueErr != nil {
				return ValueErr
			}
			stringValue := fmt.Sprintf("%s", value)

			if strings.HasSuffix(stringKey, "|id") {
				if !contains(doneMsg, stringValue) {
					// Remember we already added this message to the view
					doneMsg = append(doneMsg, stringKey)

					// Get all Data
					senderItem, senderErr := txn.Get([]byte(strings.Replace(stringKey, "|id", "|sender", -1)))
					if senderErr != nil {
						return errors.WithMessage(senderErr, "Key: "+strings.Replace(stringKey, "|id", "|sender", -1))
					}

					senderValue, senderValueErr := senderItem.Value()
					if senderValueErr != nil {
						return senderValueErr
					}
					sender := fmt.Sprintf("%s", senderValue)

					messageItem, messageErr := txn.Get([]byte(strings.Replace(stringKey, "|id", "|messageString", -1)))
					if messageErr != nil {
						return errors.WithMessage(messageErr, "Key: "+strings.Replace(stringKey, "|id", "|messageString", -1))
					}

					messageValue, messageValueErr := messageItem.Value()
					if messageValueErr != nil {
						return messageValueErr
					}
					msg := fmt.Sprintf("%s", messageValue)

					timestampItem, timestampErr := txn.Get([]byte(strings.Replace(stringKey, "|id", "|timestamp", -1)))
					if timestampErr != nil {
						return errors.WithMessage(timestampErr, "Key: "+strings.Replace(stringKey, "|id", "|timestamp", -1))
					}

					timestampValue, timestampValueErr := timestampItem.Value()
					if timestampValueErr != nil {
						return timestampValueErr
					}
					timestamp := fmt.Sprintf("%s", timestampValue)
					timestampInt, ConvErr := strconv.ParseInt(timestamp, 10, 64)
					if ConvErr != nil {
						return ConvErr
					}

					m.MessageListLayout.TriggerMessage(msg, sender, timestampInt)
				}
			}

			if strings.HasSuffix(stringKey, "|sender") {
				if !contains(doneMsg, stringValue) {
					// Remember we already added this message to the view
					idItem, idErr := txn.Get([]byte(strings.Replace(stringKey, "|sender", "|id", -1)))
					if idErr != nil {
						return errors.WithMessage(idErr, "Key: "+strings.Replace(stringKey, "|sender", "|id", -1))
					}
					idValue, idValueErr := idItem.Value()
					if idValueErr != nil {
						return idValueErr
					}
					id := fmt.Sprintf("%s", idValue)
					doneMsg = append(doneMsg, id)

					// Get all Data
					sender := stringValue

					messageItem, messageErr := txn.Get([]byte(strings.Replace(stringKey, "|sender", "|messageString", -1)))
					if messageErr != nil {
						return errors.WithMessage(messageErr, "Key: "+strings.Replace(stringKey, "|sender", "|messageString", -1))
					}

					messageValue, messageValueErr := messageItem.Value()
					if messageValueErr != nil {
						return messageValueErr
					}
					msg := fmt.Sprintf("%s", messageValue)

					timestampItem, timestampErr := txn.Get([]byte(strings.Replace(stringKey, "|sender", "|timestamp", -1)))
					if timestampErr != nil {
						return errors.WithMessage(timestampErr, "Key: "+strings.Replace(stringKey, "|sender", "|timestamp", -1))
					}

					timestampValue, timestampValueErr := timestampItem.Value()
					if timestampValueErr != nil {
						return timestampValueErr
					}
					timestamp := fmt.Sprintf("%s", timestampValue)

					timestampInt, ConvErr := strconv.ParseInt(timestamp, 10, 64)
					if ConvErr != nil {
						return ConvErr
					}

					m.MessageListLayout.TriggerMessage(msg, sender, timestampInt)
				}
			}

			if strings.HasSuffix(stringKey, "|messageString") {
				if !contains(doneMsg, stringValue) {
					// Remember we already added this message to the view
					idItem, idErr := txn.Get([]byte(strings.Replace(stringKey, "|messageString", "|id", -1)))
					if idErr != nil {
						return errors.WithMessage(idErr, "Key: "+strings.Replace(stringKey, "|messageString", "|id", -1))
					}

					idValue, idValueErr := idItem.Value()
					if idValueErr != nil {
						return idValueErr
					}
					id := fmt.Sprintf("%s", idValue)
					doneMsg = append(doneMsg, id)

					// Get all Data
					senderItem, senderErr := txn.Get([]byte(strings.Replace(stringKey, "|messageString", "|sender", -1)))
					if senderErr != nil {
						return errors.WithMessage(senderErr, "Key: "+strings.Replace(stringKey, "|messageString", "|sender", -1))
					}

					senderValue, senderValueErr := senderItem.Value()
					if senderValueErr != nil {
						return senderValueErr
					}
					sender := fmt.Sprintf("%s", senderValue)

					msg := stringValue

					timestampItem, timestampErr := txn.Get([]byte(strings.Replace(stringKey, "|messageString", "|timestamp", -1)))
					if timestampErr != nil {
						return errors.WithMessage(timestampErr, "Key: "+strings.Replace(stringKey, "|messageString", "|timestamp", -1))
					}

					timestampValue, timestampValueErr := timestampItem.Value()
					if timestampValueErr != nil {
						return timestampValueErr
					}
					timestamp := fmt.Sprintf("%s", timestampValue)

					timestampInt, ConvErr := strconv.ParseInt(timestamp, 10, 64)
					if ConvErr != nil {
						return ConvErr
					}

					m.MessageListLayout.TriggerMessage(msg, sender, timestampInt)
				}
			}

			if strings.HasSuffix(stringKey, "|timestamp") {
				if !contains(doneMsg, stringValue) {
					// Remember we already added this message to the view
					idItem, idErr := txn.Get([]byte(strings.Replace(stringKey, "|timestamp", "|id", -1)))
					if idErr != nil {
						return errors.WithMessage(idErr, "Key: "+strings.Replace(stringKey, "|timestamp", "|id", -1))
					}
					idValue, idValueErr := idItem.Value()
					if idValueErr != nil {
						return idValueErr
					}
					id := fmt.Sprintf("%s", idValue)
					doneMsg = append(doneMsg, id)

					// Get all Data
					senderItem, senderErr := txn.Get([]byte(strings.Replace(stringKey, "|timestamp", "|sender", -1)))
					if senderErr != nil {
						return errors.WithMessage(senderErr, "Key: "+strings.Replace(stringKey, "|timestamp", "|sender", -1))
					}
					senderValue, senderValueErr := senderItem.Value()
					if senderValueErr != nil {
						return senderValueErr
					}
					sender := fmt.Sprintf("%s", senderValue)

					messageItem, messageErr := txn.Get([]byte(strings.Replace(stringKey, "|timestamp", "|messageString", -1)))
					if messageErr != nil {
						return errors.WithMessage(messageErr, "Key: "+strings.Replace(stringKey, "|timestamp", "|messageString", -1))
					}
					messageValue, messageValueErr := messageItem.Value()
					if messageValueErr != nil {
						return messageValueErr
					}
					msg := fmt.Sprintf("%s", messageValue)

					timestamp := stringValue

					timestampInt, ConvErr := strconv.ParseInt(timestamp, 10, 64)
					if ConvErr != nil {
						return ConvErr
					}

					m.MessageListLayout.TriggerMessage(msg, sender, timestampInt)
				}
			}
		}

		return nil
	})
	if DBerr != nil {
		fmt.Println("DBERR: ", DBerr)
		err = DBerr
		return
	}

	/*if barAtBottom {
		bar.SetValue(bar.Maximum())
	}*/

	return
}
