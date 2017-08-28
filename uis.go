package main

import (
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/golang-commonmark/markdown"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"github.com/tidwall/buntdb"
	"log"
	"strings"
	"sync"
	"time"
)

var username string
var password string

func DoLogin(username, password, homeserverURL, userID, accessToken string, localLog *log.Logger, results chan<- *gomatrix.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	var cli *gomatrix.Client
	if accessToken != "" && homeserverURL != "" && userID != "" {
		var cliErr error
		if strings.HasPrefix(homeserverURL, "https://") {
			cli, cliErr = matrix.GetClient(homeserverURL, userID, accessToken)
		} else if strings.HasPrefix(homeserverURL, "http://") {
			cli, cliErr = matrix.GetClient(homeserverURL, userID, accessToken)
		} else {
			cli, cliErr = matrix.GetClient("https://"+homeserverURL, userID, accessToken)
		}
		if cliErr != nil {
			localLog.Println(cliErr)
		}
		cli.SetCredentials(userID, accessToken)
	} else {
		var err error
		cli, err = matrix.LoginUser(username, password)
		if err != nil {
			localLog.Println(err)
		}
	}

	results <- cli
}

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetObjectName("LoginWrapper")
	widget.SetStyleSheet("QWidget#LoginWrapper { border: 0px; };")
	topLayout := widgets.NewQVBoxLayout()

	formWidget := widgets.NewQWidget(nil, 0)
	formWrapper := widgets.NewQHBoxLayout()
	formWidget.SetMinimumSize2(350, 200)

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

	loginButton.ConnectClicked(func(checked bool) {
		//TODO register enter and show loader or so

		var wg sync.WaitGroup

		if username != "" && password != "" {
			localLog.Println("Starting Login Sequenze in background")
			results := make(chan *gomatrix.Client)

			wg.Add(1)
			go DoLogin(username, password, "", "", "", localLog, results, &wg)

			go func() {
				wg.Wait()      // wait for each execTask to return
				close(results) // then close the results channel
			}()

			//Show MainUI
			for result := range results {
				//TODO Don't switch screen on wrong login data.
				mainUI := NewMainUI(windowWidth, windowHeight, result)
				mainUI.SetMinimumSize2(windowWidth, windowHeight)
				window.SetCentralWidget(mainUI)
			}
		} else {
			localLog.Println("Username and/or password is empty. Do Nothing.")
		}
	})

	widget.SetWindowTitle("Morpheus - Login")

	return widget
}

//NewMainUI initializes the login Screen
//func NewMainUI(windowWidth, windowHeight int, cli *gomatrix.Client) *widgets.QWidget {
func NewMainUI(windowWidth, windowHeight int, cli *gomatrix.Client) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	var mainWidget = loader.Load(file, widget)
	file.Close()

	mainWidget.SetMinimumSize2(windowWidth, windowHeight)
	mainWidget.Resize2(windowWidth, windowHeight)
	mainWidget.SetGeometry2(0, 0, windowWidth, windowHeight)

	widget.SetMinimumSize2(windowWidth, windowHeight)
	widget.SetGeometry2(0, 0, windowWidth, windowHeight)

	chatWidget := widgets.NewQWidgetFromPointer(widget.FindChild("ChatWidget", core.Qt__FindChildrenRecursively).Pointer())
	chatWidget.SetMinimumSize2(windowWidth, windowHeight)
	chatWidget.SetGeometry2(0, 0, windowWidth, windowHeight)

	mainWidget.SetWindowTitle("Morpheus - MatrixHQ")
	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(mainWidget, 1, core.Qt__AlignTop|core.Qt__AlignLeft)
	widget.SetLayout(layout)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)

	window.ConnectResizeEvent(func(event *gui.QResizeEvent) {

		widget.SetGeometry2(0, 0, event.Size().Width(), event.Size().Height())
		mainWidget.SetGeometry2(0, 0, event.Size().Width(), event.Size().Height())
		chatWidget.SetGeometry2(0, 0, event.Size().Width(), event.Size().Height())

		widget.Resize(event.Size())
		mainWidget.Resize(event.Size())
		chatWidget.Resize(event.Size())
	})

	//Set Avatar
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	avatarLogo.SetPixmap(matrix.GetOwnUserAvatar(cli))

	//Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(checked bool) {
		//TODO register enter and show loader or so
		localLog.Println("Starting Logout Sequenze in background")
		var wg sync.WaitGroup
		results := make(chan bool)

		wg.Add(1)
		go func(cli *gomatrix.Client, localLog *log.Logger, results chan<- bool) {
			defer wg.Done()
			_, err := cli.Logout()
			if err != nil {
				localLog.Println(err)
				results <- false
			} else {
				cli.ClearCredentials()
				//Flush complete DB
				db.View(func(tx *buntdb.Tx) error {
					var QueryErr error
					QueryErr = tx.DeleteAll()
					if QueryErr != nil {
						return QueryErr
					}
					return nil
				})
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
				loginUI := NewLoginUI(windowWidth, windowHeight)
				loginUI.SetMinimumSize2(windowWidth, windowHeight)
				window.SetCentralWidget(loginUI)
			}
		}
	})

	//Start Syncer!
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	customStore := gomatrix.NewInMemoryStore()
	cli.Store = customStore
	syncer.Store = customStore
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		localLog.Println("message:", ev.Content)
	})

	// Start Non-blocking sync
	localLog.Println("Syncing now")
	go func() {
		for {
			localLog.Println("sync")
			if e := cli.Sync(); e != nil {
				localLog.Println("Fatal Sync() error")
				time.Sleep(10 * time.Second)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	messageInput := widgets.NewQLineEditFromPointer(widget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
	var message string
	window.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
			mardownMessage := md.RenderToString([]byte(message))
			if mardownMessage == message {
				cli.SendText("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", message)
			} else {
				cli.SendMessageEvent("!zTIXGmDjyRcAqbrWab:matrix.ffslfl.net", "m.room.message", matrix.HTMLMessage{"m.text", message, mardownMessage, "org.matrix.custom.html"})
			}
			messageInput.Clear()
		} else {
			messageInput.KeyPressEventDefault(ev)
		}
	})
	messageInput.ConnectTextChanged(func(value string) {
		message = value
	})

	localLog.Println("Started Syncing")

	return widget
}
