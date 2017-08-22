package main

import (
	"fmt"

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
		localLog.Println("Starting Login Sequenze")
		cli, err := matrix.LoginUser(username, password)
		if err != nil {
			localLog.Println(err)
		}
		MainUI := NewMainUI(windowWidth, windowHeight, cli)

		//Show MainUI
		window.SetCentralWidget(MainUI)
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
	DisplayNameResp, DisplayNameErr := cli.GetUserDisplayName(username)
	if DisplayNameErr != nil {
		localLog.Println(DisplayNameErr)
	}
	usernameLabel.SetText(fmt.Sprint(DisplayNameResp.DisplayName))

	// Set Avatar
	avatarLogo.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	avatarLogo.SetPixmap(cli.GetOwnUserAvatar())

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(mainWidget, 0, 0)
	widget.SetLayout(layout)

	return widget
}
