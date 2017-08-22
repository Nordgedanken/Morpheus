package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var username string
var password string

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()
	widget.SetLayout(layout)

	// UsernameInput
	usernameInput := widgets.NewQLineEdit(nil)
	usernameInput.SetPlaceholderText("Insert MXID")
	layout.AddWidget(usernameInput, 0, 0)

	// PasswordInput
	passwordInput := widgets.NewQLineEdit(nil)
	passwordInput.SetPlaceholderText("Insert password")
	passwordInput.SetEchoMode(widgets.QLineEdit__Password)
	layout.AddWidget(passwordInput, 0, 0)

	// UsernameInput-Label
	usernameLabel := widgets.NewQLabel(nil, 0)
	usernameLabel.SetText("Username: ")
	usernameLabel.SetBuddy(usernameInput)
	layout.AddWidget(usernameLabel, 0, 0)

	// PasswordInput-Label
	passwordLabel := widgets.NewQLabel(nil, 0)
	passwordLabel.SetText("Password: ")
	passwordLabel.SetBuddy(passwordInput)
	layout.AddWidget(passwordLabel, 0, 0)

	// loginButton
	loginButton := widgets.NewQPushButton2("LOGIN", nil)
	layout.AddWidget(loginButton, 0, 0)

	usernameInput.ConnectTextChanged(func(value string) {
		username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		password = value
	})

	loginButton.ConnectClicked(func(checked bool) {
		//TODO register enter and show loader or so
		localLog.Println("Starting Login Sequenze in background")
		var wg sync.WaitGroup
		results := make(chan *matrix.Client)

		wg.Add(1)
		go func(username, password string, localLog *log.Logger, results chan<- *matrix.Client) {
			defer wg.Done()
			cli, err := matrix.LoginUser(username, password)
			if err != nil {
				localLog.Println(err)
			}

			results <- cli

		}(username, password, localLog, results)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show MainUI
		for result := range results {
			MainUI := NewMainUI(windowWidth, windowHeight, result)
			window.SetCentralWidget(MainUI)
		}

	})

	widget.SetWindowTitle("Neo - Login")

	return widget
}

//NewMainUI initializes the login Screen
func NewMainUI(windowWidth, windowHeight int, cli *matrix.Client) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/Main.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, widget)
	file.Close()

	var (
		usernameLabel = widgets.NewQLabelFromPointer(widget.FindChild("UsernameLabel", core.Qt__FindChildrenRecursively).Pointer())
		mxidLabel     = widgets.NewQLabelFromPointer(widget.FindChild("MXIDLabel", core.Qt__FindChildrenRecursively).Pointer())
		avatarLogo    = widgets.NewQLabelFromPointer(widget.FindChild("AvatarLabel", core.Qt__FindChildrenRecursively).Pointer())
	)

	// Set MXID Label
	mxidLabel.SetText(fmt.Sprint(username))

	// Set Dispalyname Label
	displayName, displayNameErr := cli.GetUserDisplayName(username)
	if displayNameErr != nil {
		localLog.Println(displayNameErr)
	}
	usernameLabel.SetText(fmt.Sprint(displayName))

	// Set Avatar
	avatarLogo.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	avatarLogo.SetPixmap(cli.GetOwnUserAvatar())

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(mainWidget, 0, 0)
	widget.SetLayout(layout)

	return widget
}
