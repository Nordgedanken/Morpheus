package main

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/ui"
	"github.com/Nordgedanken/dugong"
	"github.com/matrix-org/gomatrix"
	"github.com/shibukawa/configdir"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var window *widgets.QMainWindow
var mainUIStruct *ui.MainUI
var loginUIStruct *ui.LoginUI

func main() {
	runtime.GOMAXPROCS(128)

	// Init Logs
	configDirs := configdir.New("Nordgedanken", "Morpheus")

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05.000000",
		DisableColors:    false,
		DisableTimestamp: true,
		DisableSorting:   true,
		QuoteEmptyFields: true,
	})

	log.AddHook(dugong.NewFSHook(
		filepath.Join(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/", "info.log"),
		filepath.Join(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/", "warn.log"),
		filepath.Join(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/log/", "error.log"),
		&log.TextFormatter{
			TimestampFormat:  "2006-01-02 15:04:05.000000",
			DisableColors:    true,
			DisableTimestamp: false,
			DisableSorting:   false,
		}, &dugong.DailyRotationSchedule{GZip: false},
	))

	log.Infoln("Starting Morpheus")

	app := widgets.NewQApplication(len(os.Args), os.Args)

	app.SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)
	app.SetApplicationName("Morpheus")
	app.SetApplicationVersion("0.0.1")
	appIcon := gui.NewQIcon5(":/qml/resources/logos/MorpheusBig.png")
	app.SetWindowIcon(appIcon)

	window = widgets.NewQMainWindow(nil, 0)
	/*window.SetWindowFlags(core.Qt__FramelessWindowHint)
	var dragPositionX int
	var dragPositionY int
	window.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		if event.Button() == core.Qt__LeftButton {
			dragPositionX = event.GlobalX() - window.FrameGeometry().TopLeft().X()
			dragPositionY = event.GlobalY() - window.FrameGeometry().TopLeft().Y()
			event.Accept()
		}
	})

	window.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		if event.Button() == core.Qt__LeftButton {
			window.Move2(event.GlobalX()-dragPositionX, event.GlobalY()-dragPositionY)
			event.Accept()
		}
	})*/

	windowHeight := 600
	windowWidth := 950

	desktopApp := widgets.QApplication_Desktop()
	primaryScreen := desktopApp.PrimaryScreen()
	screen := desktopApp.Screen(primaryScreen)
	windowX := (screen.Width() - windowHeight) / 2
	windowY := (screen.Height() - windowWidth) / 2

	window.Resize2(windowWidth, windowHeight)
	window.Show()

	window.Move2(windowX, windowY)

	accessToken, homeserverURL, userID, UserCacheErr := matrix.GetUserDataFromCache()
	if UserCacheErr != nil {
		log.Debug(UserCacheErr)
	}

	if accessToken != "" && homeserverURL != "" && userID != "" {
		var wg sync.WaitGroup
		log.Infoln("Starting Auto Login Sequenze in background")
		results := make(chan *gomatrix.Client)

		wg.Add(1)
		go matrix.DoLogin("", "", homeserverURL, userID, accessToken, results, &wg)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show MainUI
		for result := range results {
			mainUIStruct = ui.NewMainUIStruct(windowWidth, windowHeight, window)
			mainUIStruct.SetCli(result)
			mainUIErr := mainUIStruct.NewUI()
			if mainUIErr != nil {
				log.Errorln("mainUI: ", mainUIErr)
				return
			}
			scalar.ReqAndSaveScalarToken(MainUIStruct.GetCli())

			mainUIStruct.GetWidget().Resize2(windowWidth, windowHeight)
			window.SetCentralWidget(mainUIStruct.GetWidget())
		}
	} else {
		//Show loginUI
		loginUIStruct = ui.NewLoginUIStruct(windowWidth, windowHeight, window)
		loginUIErr := loginUIStruct.NewUI()
		if loginUIErr != nil {
			log.Errorln("Login Err: ", loginUIErr)
			return
		}

		loginUIStruct.GetWidget().Resize2(windowWidth, windowHeight)
		window.SetCentralWidget(loginUIStruct.GetWidget())
	}

	window.Resize2(windowWidth, windowHeight)

	window.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		log.Infoln("Stopping Morpheus")
		if cleanup() {
			event.Accept()
		} else {
			event.Ignore()
		}
	})

	//enter the main event loop
	_ = widgets.QApplication_Exec()
}

func cleanup() bool {
	log.Infoln("cleanup")

	if mainUIStruct != nil {
		if mainUIStruct.GetCli() != nil {
			log.Infoln("Stop Sync")
			mainUIStruct.GetCli().StopSync()
		}
	}

	if loginUIStruct != nil {
		if loginUIStruct.GetCli() != nil {
			log.Infoln("Stop Sync")
			loginUIStruct.GetCli().StopSync()
		}
	}

	UserDB, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return false
	}

	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return false
	}

	UserDB.Close()
	CacheDB.Close()

	return true
}
