package main

import (
	"fmt"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var username string
var password string

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/login.ui")

	file.Open(core.QIODevice__ReadOnly)
	loginWidget := loader.Load(file, widget)
	file.Close()

	var (
		inputUsername = widgets.NewQLineEditFromPointer(widget.FindChild("UsernameInput", core.Qt__FindChildrenRecursively).Pointer())
		inputPassword = widgets.NewQLineEditFromPointer(widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())
		SubmitButton  = widgets.NewQPushButtonFromPointer(widget.FindChild("loginButton", core.Qt__FindChildrenRecursively).Pointer())
	)

	inputUsername.ConnectTextChanged(func(value string) {
		username = value
	})

	inputPassword.ConnectTextChanged(func(value string) {
		password = value
	})

	SubmitButton.ConnectClicked(func(checked bool) {
		localLog.Println("Starting Login Sequenze")
		cli, err := matrix.LoginUser(username, password)
		if err != nil {
			localLog.Fatalln(err)
		}
		MainUI := NewMainUI(windowWidth, windowHeight, cli)
		MainUI.SetMinimumSize2(windowWidth, windowHeight)

		//Show MainUI
		widget.Close()
		MainUI.Show()
	})

	loginWidget.SetMinimumSize2(windowWidth, windowHeight)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(loginWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Neo - Login")

	return widget
}

//NewMainUI initializes the login Screen
func NewMainUI(windowWidth, windowHeight int, cli *gomatrix.Client) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/Main.ui")

	file.Open(core.QIODevice__ReadOnly)
	loginWidget := loader.Load(file, widget)
	file.Close()

	var (
		usernameLabel = widgets.NewQLabelFromPointer(widget.FindChild("UsernameLabel", core.Qt__FindChildrenRecursively).Pointer())
		mxidLabel     = widgets.NewQLabelFromPointer(widget.FindChild("MXIDLabel", core.Qt__FindChildrenRecursively).Pointer())
		//avatarLogo    = widgets.NewQGraphicsViewFromPointer(widget.FindChild("loginButton", core.Qt__FindChildrenRecursively).Pointer())
	)

	// Set MXID Label
	mxidLabel.SetText(fmt.Sprint(username))

	// Set Dispalyname Label
	DisplayNameResp, DisplayNameErr := matrix.GetUserDisplayName(username, cli)
	if DisplayNameErr != nil {
		localLog.Println(DisplayNameErr)
	}
	usernameLabel.SetText(fmt.Sprint(DisplayNameResp.DisplayName))

	//TODO remove after testing
	avatarURL, avatarErr := cli.GetAvatarURL()
	if avatarErr != nil {
		localLog.Println(avatarErr)
	}
	localLog.Println("Avatar: " + avatarURL)

	loginWidget.SetMinimumSize2(windowWidth, windowHeight)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(loginWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Neo - Login")

	return widget
}
