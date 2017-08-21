package main

import (
	"log"
	"os"

	"github.com/Nordgedanken/Neo/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var localLog *log.Logger

func main() {
	localLog = util.Logger()
	util.StartFileLog(localLog)

	widgets.NewQApplication(len(os.Args), os.Args)

	desktopApp := widgets.QApplication_Desktop()
	primaryScreen := desktopApp.PrimaryScreen()
	screen := desktopApp.Screen(primaryScreen)
	windowWidth := (screen.Width() / 2)
	windowHeight := (screen.Height() / 2)

	loginUI := NewLoginUI(windowWidth, windowHeight)
	loginUI.SetMinimumSize2(windowWidth, windowHeight)

	//Show loginUI
	loginUI.Show()

	//enter the main event loop
	widgets.QApplication_Exec()
}

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
