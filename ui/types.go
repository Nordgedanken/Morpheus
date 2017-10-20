package ui

import (
	"log"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/widgets"
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

	widget            *widgets.QWidget
	RoomAvatar        *widgets.QLabel
	RoomTitle         *widgets.QLabel
	RoomTopic         *widgets.QLabel
	MainWidget        *widgets.QWidget
	MessageListLayout *QVBoxLayoutWithTriggerSlot

	window      *widgets.QMainWindow
	syncer      *gomatrix.DefaultSyncer
	storage     *gomatrix.InMemoryStore
	rooms       map[string]*matrix.Room
	currentRoom string
}

func (m *MainUI) SetCurrentRoom(RoomID string) {
	m.currentRoom = RoomID
}

// LoginUI holds information about the LoginUI
type LoginUI struct {
	logger
	config
	widget *widgets.QWidget
	window *widgets.QMainWindow
}
