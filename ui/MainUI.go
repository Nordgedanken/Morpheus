package ui

import (
	"github.com/Nordgedanken/Morpheus/elements"
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/util"
	"github.com/matrix-org/gomatrix"
	"github.com/rhinoman/go-commonmark"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"github.com/tidwall/buntdb"
	"log"
	"sync"
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
	}
	return
}

// NewMainUIStructWithExistingConfig gives you a MainUI struct with prefilled data and data from a previous Config
func NewMainUIStructWithExistingConfig(configStruct config, window *widgets.QMainWindow) (mainUIStruct MainUI) {
	mainUIStruct = MainUI{
		config: configStruct,
		window: window,
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
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	var mainWidget = loader.Load(file, widget)
	file.Close()

	InitDataErr := matrix.InitData(m.cli)
	if InitDataErr != nil {
		err = InitDataErr
		return
	}

	messageScrollArea := widgets.NewQScrollAreaFromPointer(widget.FindChild("messageScroll", core.Qt__FindChildrenRecursively).Pointer())
	messagesScrollAreaContent := widgets.NewQWidgetFromPointer(widget.FindChild("messagesScrollAreaContent", core.Qt__FindChildrenRecursively).Pointer())

	mainWidget.SetWindowTitle("Morpheus - MatrixHQ")

	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(mainWidget, 1, core.Qt__AlignTop|core.Qt__AlignLeft)
	widget.SetLayout(layout)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)

	m.window.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		widget.Resize2(event.Size().Width(), event.Size().Height())
		event.Accept()
	})

	widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		mainWidget.Resize2(event.Size().Width(), event.Size().Height())
		event.Accept()
	})

	messageScrollArea.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		messageScrollArea.Resize(event.Size())
		event.Accept()
	})

	//Set Avatar
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatar, AvatarErr := matrix.GetOwnUserAvatar(m.cli)
	if AvatarErr != nil {
		err = AvatarErr
		return
	}
	avatarLogo.SetPixmap(avatar)

	//Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		LogoutErr := m.logout(widget, messageScrollArea)
		if LogoutErr != nil {
			err = LogoutErr
			return
		}
	})

	// Init Message View
	messageListLayout := elements.NewMessageList(messageScrollArea, messagesScrollAreaContent)

	messageListLayout.ConnectTriggerMessage(func(messageBody, sender string) {
		NewMessageErr := messageListLayout.NewMessage(messageBody, m.cli, sender, messageScrollArea)
		if NewMessageErr != nil {
			err = NewMessageErr
			return
		}
	})
	messageScrollArea.SetWidgetResizable(true)
	messageScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	messageScrollArea.SetContentsMargins(0, 0, 0, 0)

	m.startSync(messageListLayout)

	var message string
	messageInput := widgets.NewQLineEditFromPointer(widget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
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

	m.widget = widget
	return
}

func (m *MainUI) sendMessage(message string) (err error) {
	mardownMessage := commonmark.Md2Html(message, 0)
	if mardownMessage == message {
		_, SendErr := m.cli.SendText("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", message)
		if SendErr != nil {
			err = SendErr
			return
		}
	} else {
		_, SendErr := m.cli.SendMessageEvent("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", "m.room.message", matrix.HTMLMessage{MsgType: "m.text", Body: message, FormattedBody: mardownMessage, Format: "org.matrix.custom.html"})
		if SendErr != nil {
			err = SendErr
			return
		}
	}
	return
}

func (m *MainUI) logout(widget *widgets.QWidget, messageScrollArea *widgets.QScrollArea) (err error) {
	//TODO register enter and show loader or so
	m.localLog.Println("Starting Logout Sequenze in background")
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

		db, DBOpenErr := matrix.OpenDB()
		if DBOpenErr != nil {
			localLog.Fatalln(DBOpenErr)
		}
		defer db.Close()

		//Flush complete DB
		DBErr := db.Update(func(tx *buntdb.Tx) error {
			QueryErr := tx.DeleteAll()
			return QueryErr
		})
		if DBErr != nil {
			localLog.Panicln(DBErr)
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

func (m *MainUI) startSync(messageListLayout *elements.QVBoxLayoutWithTriggerSlot) (err error) {
	//Start Syncer!
	m.syncer = m.cli.Syncer.(*gomatrix.DefaultSyncer)
	customStore := gomatrix.NewInMemoryStore()
	m.cli.Store = customStore
	m.syncer.Store = customStore

	m.syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		msg, _ := ev.Body()
		room := ev.RoomID
		sender := ev.Sender
		id := ev.ID
		timestamp := ev.Timestamp
		go matrix.CacheMessageEvents(id, sender, room, msg, timestamp)
		messageListLayout.TriggerMessage(msg, sender)
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
