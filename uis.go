package main

import (
	"log"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/Nordgedanken/Neo/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var localLog *log.Logger

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	var username string
	var password string
	localLog = util.Logger()
	util.StartFileLog(localLog)

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
			localLog.Panicln(err)
		}

		//TODO remove after testing
		avatarURL, avatarErr := cli.GetAvatarURL()
		if avatarErr != nil {
			localLog.Panicln(avatarErr)
		}
		localLog.Println(avatarURL)
	})

	loginWidget.SetMinimumSize2(windowWidth, windowHeight)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(loginWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Neo - Login")

	return widget
}
