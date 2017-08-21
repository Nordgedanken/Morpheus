package main

import (
	"log"
	"os"

	"github.com/Nordgedanken/Neo/util"
	"github.com/therecipe/qt/widgets"
)

var localLog *log.Logger

func main() {
	localLog = util.Logger()
	localLog = util.StartFileLog(localLog)
	localLog.Println("Starting Neo")

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
	localLog.Println("Finished Startup Neo")
}
