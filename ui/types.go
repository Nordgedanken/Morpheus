package ui

import (
	"github.com/Nordgedanken/Morpheus/matrix/globalTypes"
	"github.com/Nordgedanken/Morpheus/matrix/syncer"
	"github.com/therecipe/qt/widgets"
)

// MainUI holds information about the MainUI
type MainUI struct {
	globalTypes.Config

	widget            *widgets.QWidget
	RoomAvatar        *widgets.QLabel
	RoomTitle         *widgets.QLabel
	RoomTopic         *widgets.QLabel
	MainWidget        *widgets.QWidget
	messageScrollArea *widgets.QScrollArea
	roomScrollArea    *widgets.QScrollArea

	window  *widgets.QMainWindow
	storage *syncer.MorpheusStore
}

// SetCurrentRoom sets the new room ID of the MainUI
func (m *MainUI) SetCurrentRoom(RoomID string) {
	m.CurrentRoom = RoomID
}

// LoginUI holds information about the LoginUI
type LoginUI struct {
	globalTypes.Config

	LoginWidget *widgets.QWidget
	widget      *widgets.QWidget
	window      *widgets.QMainWindow
}
