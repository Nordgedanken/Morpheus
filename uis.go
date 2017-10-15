package main

import (
	"github.com/Nordgedanken/Morpheus/elements"
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/matrix-org/gomatrix"
	"github.com/rhinoman/go-commonmark"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"github.com/tidwall/buntdb"
	"log"
	"sync"
	"time"
)

var username string
var password string

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int, window *widgets.QMainWindow) (respWidget *widgets.QWidget, err error) {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetObjectName("LoginWrapper")
	widget.SetStyleSheet("QWidget#LoginWrapper { border: 0px; };")
	topLayout := widgets.NewQVBoxLayout()

	formWidget := widgets.NewQWidget(nil, 0)
	formWrapper := widgets.NewQHBoxLayout()

	formLayout := widgets.NewQVBoxLayout()
	formLayout.SetSpacing(20)
	formLayout.SetContentsMargins(0, 0, 0, 30)
	formWidget.SetLayout(formLayout)

	formWrapper.AddStretch(1)
	formWrapper.AddWidget(formWidget, 0, 0)
	formWrapper.AddStretch(1)

	// UsernameInput
	usernameInput := widgets.NewQLineEdit(nil)
	usernameInput.SetPlaceholderText("Insert MXID")

	usernameLayout := widgets.NewQHBoxLayout()
	usernameLayout.AddWidget(usernameInput, 0, core.Qt__AlignVCenter)

	// PasswordInput
	passwordInput := widgets.NewQLineEdit(nil)
	passwordInput.SetPlaceholderText("Insert password")
	passwordInput.SetEchoMode(widgets.QLineEdit__Password)

	passwordLayout := widgets.NewQHBoxLayout()
	passwordLayout.AddWidget(passwordInput, 0, core.Qt__AlignVCenter)

	formLayout.AddLayout(usernameLayout, 0)
	formLayout.AddLayout(passwordLayout, 0)

	// loginButton
	buttonLayout := widgets.NewQHBoxLayout()
	buttonLayout.SetSpacing(0)
	buttonLayout.SetContentsMargins(0, 0, 0, 30)

	loginButton := widgets.NewQPushButton2("LOGIN", nil)
	loginButton.SetMinimumSize2(350, 65)

	buttonLayout.AddStretch(1)
	buttonLayout.AddWidget(loginButton, 0, 0)
	buttonLayout.AddStretch(1)

	topLayout.AddStretch(1)
	topLayout.AddLayout(formWrapper, 0)
	topLayout.AddStretch(1)
	topLayout.AddLayout(buttonLayout, 0)
	topLayout.AddStretch(1)

	widget.SetLayout(topLayout)

	usernameInput.ConnectTextChanged(func(value string) {
		username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		password = value
	})

	loginButton.ConnectClicked(func(_ bool) {
		//TODO register enter and show loader or so

		var wg sync.WaitGroup

		if username != "" && password != "" {
			localLog.Println("Starting Login Sequenze in background")
			results := make(chan *gomatrix.Client)

			wg.Add(1)
			go matrix.DoLogin(username, password, "", "", "", localLog, results, &wg)

			go func() {
				wg.Wait()      // wait for each execTask to return
				close(results) // then close the results channel
			}()

			//Show MainUI
			for result := range results {
				//TODO Don't switch screen on wrong login data.
				mainUI, mainUIErr := NewMainUI(windowWidth, windowHeight, result, window)
				if mainUIErr != nil {
					err = mainUIErr
					return
				}
				window.SetCentralWidget(mainUI)
				window.Resize(widget.Size())
			}
		} else {
			localLog.Println("Username and/or password is empty. Do Nothing.")
		}
	})

	widget.SetWindowTitle("Morpheus - Login")

	respWidget = widget
	return
}

//NewMainUI initializes the login Screen
//func NewMainUI(windowWidth, windowHeight int, cli *gomatrix.Client) *widgets.QWidget {
func NewMainUI(windowWidth, windowHeight int, cli *gomatrix.Client, window *widgets.QMainWindow) (respWidget *widgets.QWidget, err error) {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	var mainWidget = loader.Load(file, widget)
	file.Close()

	InitDataErr := matrix.InitData(cli)
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

	window.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		widget.Resize2(event.Size().Width(), event.Size().Height())

	})

	widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		mainWidget.Resize2(event.Size().Width(), event.Size().Height())

	})

	messageScrollArea.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		messageScrollArea.Resize(event.Size())
	})

	//Set Avatar
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatar, AvatarErr := matrix.GetOwnUserAvatar(cli)
	if AvatarErr != nil {
		err = AvatarErr
		return
	}
	avatarLogo.SetPixmap(avatar)

	//Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		//TODO register enter and show loader or so
		localLog.Println("Starting Logout Sequenze in background")
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
		}(cli, localLog, results)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show LoginUI
		for result := range results {
			if result {
				window.DisconnectKeyPressEvent()
				window.DisconnectResizeEvent()
				widget.DisconnectResizeEvent()
				messageScrollArea.DisconnectResizeEvent()

				loginUI, loginUIErr := NewLoginUI(windowWidth, windowHeight, window)
				if loginUIErr != nil {
					localLog.Panicln(loginUIErr)
				}
				window.SetCentralWidget(loginUI)
			}
		}
	})

	//Start Syncer!
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	customStore := gomatrix.NewInMemoryStore()
	cli.Store = customStore
	syncer.Store = customStore

	// Init Message View
	messageListLayout := elements.NewMessageList(messageScrollArea, messagesScrollAreaContent)

	messageListLayout.ConnectTriggerMessage(func(messageBody, sender string) {
		NewMessageErr := messageListLayout.NewMessage(messageBody, cli, sender, messageScrollArea)
		if NewMessageErr != nil {
			err = NewMessageErr
			return
		}
	})
	messageScrollArea.SetWidgetResizable(true)
	messageScrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	messageScrollArea.SetContentsMargins(0, 0, 0, 0)

	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		msg, _ := ev.Body()
		room := ev.RoomID
		sender := ev.Sender
		id := ev.ID
		timestamp := ev.Timestamp
		CacheErr := matrix.CacheMessageEvents(id, sender, room, msg, timestamp)
		if CacheErr != nil {
			err = CacheErr
			return
		}
		messageListLayout.TriggerMessage(msg, sender)
	})

	// Start Non-blocking sync
	go func() {
		localLog.Println("Start sync")
		for {

			if e := cli.Sync(); e != nil {
				localLog.Println("Fatal Sync() error")
				time.Sleep(5 * time.Second)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	messageInput := widgets.NewQLineEditFromPointer(widget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
	var message string
	window.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			mardownMessage := commonmark.Md2Html(message, 0)
			if mardownMessage == message {
				_, SendErr := cli.SendText("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", message)
				if SendErr != nil {
					err = SendErr
					return
				}
			} else {
				_, SendErr := cli.SendMessageEvent("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", "m.room.message", matrix.HTMLMessage{"m.text", message, mardownMessage, "org.matrix.custom.html"})
				if SendErr != nil {
					err = SendErr
					return
				}
			}
			messageInput.Clear()
		} else {
			messageInput.KeyPressEventDefault(ev)
		}
	})
	messageInput.ConnectTextChanged(func(value string) {
		message = value
	})

	respWidget = widget
	return
}
