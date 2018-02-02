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
	App               *widgets.QApplication
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

	LoginWidget     *widgets.QWidget
	widget          *widgets.QWidget
	App             *widgets.QApplication
	window          *widgets.QMainWindow
	helloMatrixResp helloMatrixResp
}

// RegUI holds information about the LoginUI
type RegUI struct {
	globalTypes.Config

	RegWidget       *widgets.QWidget
	widget          *widgets.QWidget
	window          *widgets.QMainWindow
	helloMatrixResp helloMatrixResp
}

type helloMatrixResp []struct {
	Hostname             string `json:"hostname"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	Category             string `json:"category"`
	Location             string `json:"location"`
	OnlineSince          int64  `json:"online_since"`
	LastResponse         int64  `json:"last_response"`
	LastResponseTime     int64  `json:"last_response_time"`
	StatusSince          string `json:"status_since"`
	LastVersions         string `json:"last_versions"`
	Measurements         int64  `json:"measurements"`
	Successful           int64  `json:"successful"`
	SumResponseTime      int64  `json:"sum_response_time"`
	MeasurementsShort    int64  `json:"measurements_short"`
	SuccessfulShort      int64  `json:"successful_short"`
	SumResponseTimeShort int64  `json:"sum_response_time_short"`
	UsersActive          int64  `json:"users_active,omitempty"`
	ServerName           string `json:"server_name"`
	ServerVersion        string `json:"server_version"`
	Grade                string `json:"grade"`
	GradeTrustIgnored    string `json:"gradeTrustIgnored"`
	HasWarnings          int64  `json:"hasWarnings"`
	PublicRoomCount      int64  `json:"public_room_count"`
}
