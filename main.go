package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/ui"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/dugong"
	"github.com/matrix-org/gomatrix"
	"github.com/shibukawa/configdir"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var window *widgets.QMainWindow

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

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	UserDB, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

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

	window.Move2(windowX, windowY)

	var accessToken string
	var homeserverURL string
	var userID string

	// Get cache
	DBErr := UserDB.View(func(txn *badger.Txn) error {
		accessTokenItem, QueryErr := txn.Get([]byte("user|accessToken"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			accessTokenByte, accessTokenErr := accessTokenItem.Value()
			accessToken = fmt.Sprintf("%s", accessTokenByte)
			if accessTokenErr != nil {
				return accessTokenErr
			}
		}

		homeserverURLItem, QueryErr := txn.Get([]byte("user|homeserverURL"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			homeserverURLByte, homeserverURLErr := homeserverURLItem.Value()
			homeserverURL = fmt.Sprintf("%s", homeserverURLByte)
			if homeserverURLErr != nil {
				return homeserverURLErr
			}
		}

		userIDItem, QueryErr := txn.Get([]byte("user|userID"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			userIDByte, userIDErr := userIDItem.Value()
			userID = fmt.Sprintf("%s", userIDByte)
			return userIDErr
		}
		return nil
	})
	if DBErr != nil {
		log.Errorln("Login: ", DBErr)
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
			//TODO Don't switch screen on wrong login data.
			MainUIStruct := ui.NewMainUIStruct(windowWidth, windowHeight, window)
			MainUIStruct.SetCli(result)
			mainUIErr := MainUIStruct.NewUI()
			if mainUIErr != nil {
				log.Errorln("mainUI: ", mainUIErr)
				return
			}
			MainUIStruct.GetWidget().Resize2(windowWidth, windowHeight)
			window.SetCentralWidget(MainUIStruct.GetWidget())
		}
	} else {
		//Show loginUI
		LoginUIStruct := ui.NewLoginUIStruct(windowWidth, windowHeight, window)
		loginUIErr := LoginUIStruct.NewUI()
		if loginUIErr != nil {
			log.Errorln("Login Err: ", loginUIErr)
			return
		}
		LoginUIStruct.GetWidget().Resize2(windowWidth, windowHeight)
		window.SetCentralWidget(LoginUIStruct.GetWidget())
	}

	window.Resize2(windowWidth, windowHeight)
	window.Show()

	//enter the main event loop
	_ = widgets.QApplication_Exec()
	defer UserDB.Close()
	defer CacheDB.Close()
	log.Infoln("Stopping Morpheus")
}

func cleanup() {
	log.Infoln("cleanup")
	UserDB, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
	}

	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
	}

	UserDB.Close()
	CacheDB.Close()
}
