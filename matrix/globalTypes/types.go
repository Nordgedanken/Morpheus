package globalTypes

import (
	"github.com/Nordgedanken/Morpheus/matrix/rooms"
	"github.com/Nordgedanken/Morpheus/ui/listLayouts"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/widgets"
)

// Config holds important reused information in the UI
type Config struct {
	Localpart string
	Password  string
	Server    string

	WindowWidth  int
	WindowHeight int

	Rooms       map[string]*rooms.Room
	CurrentRoom string

	MessageList *listLayouts.MessageList
	RoomList    *listLayouts.RoomList

	App *widgets.QApplication

	matrixClient
}

type matrixClient struct {
	Cli *gomatrix.Client
}

// GetCli returns the Matrix Client
func (mc *matrixClient) GetCli() *gomatrix.Client {
	return mc.Cli
}
