package main

import (
	"log"
	"os"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/Nordgedanken/Neo/util"
	"github.com/therecipe/qt/widgets"
)

var window *widgets.QMainWindow
var localLog *log.Logger

func main() {
	var file *os.File
	localLog = util.Logger()
	localLog, file = util.StartFileLog(localLog)
	defer file.Close()

	db := matrix.OpenDB()
	defer db.Close()

	localLog.Println("Starting Neo")

	widgets.NewQApplication(len(os.Args), os.Args)

	desktopApp := widgets.QApplication_Desktop()
	primaryScreen := desktopApp.PrimaryScreen()
	screen := desktopApp.Screen(primaryScreen)
	windowWidth := (screen.Width() / 2)
	windowHeight := (screen.Height() / 2)

	window = widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(windowWidth, windowHeight)

	loginUI := NewLoginUI(windowWidth, windowHeight)
	loginUI.Resize2(windowWidth, windowHeight)
	loginUI.SetMinimumSize2(windowWidth, windowHeight)

	//Show loginUI
	window.SetCentralWidget(loginUI)
	window.Show()

	//enter the main event loop
	widgets.QApplication_Exec()
	localLog.Println("Stopping Neo")
}
