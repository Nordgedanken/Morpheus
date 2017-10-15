package ui

import (
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/widgets"
	"log"
)

// Config holds important reused information in the UI
type config struct {
	username string
	password string

	windowWidth  int
	windowHeight int

	cli *gomatrix.Client
}

// Logger holds the initialised Logger
type logger struct {
	localLog *log.Logger
}

// MainUI holds information about the MainUI
type MainUI struct {
	logger
	config
	widget *widgets.QWidget
	window *widgets.QMainWindow
	syncer *gomatrix.DefaultSyncer
}

// LoginUI holds information about the LoginUI
type LoginUI struct {
	logger
	config
	widget *widgets.QWidget
	window *widgets.QMainWindow
}
