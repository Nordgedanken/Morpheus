package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	var username string
	var password string
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/login.ui")

	file.Open(core.QIODevice__ReadOnly)
	loginWidget := loader.Load(file, widget)
	file.Close()

	var (
		ui_inputUsername = widgets.NewQLineEditFromPointer(widget.FindChild("UsernameInput", core.Qt__FindChildrenRecursively).Pointer())
		ui_inputPassword = widgets.NewQLineEditFromPointer(widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())
		ui_SubmitButton  = widgets.NewQPushButtonFromPointer(widget.FindChild("loginButton", core.Qt__FindChildrenRecursively).Pointer())
	)

	ui_inputUsername.ConnectTextChanged(func(value string) {
		username = value
	})

	ui_inputPassword.ConnectTextChanged(func(value string) {
		password = value
	})

	ui_SubmitButton.ConnectClicked(func(checked bool) {
		localLog.Println(username + " - " + password)
	})

	loginWidget.SetMinimumSize2(windowWidth, windowHeight)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(loginWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Neo - Login")

	return widget
}
