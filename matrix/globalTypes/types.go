package globalTypes

import (
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/matrix/rooms"
	"github.com/Nordgedanken/Morpheus/matrix/syncer"
	"github.com/Nordgedanken/Morpheus/ui/listLayouts"
	"github.com/matrix-org/gomatrix"
)

// Config holds important reused information in the UI
type Config struct {
	Username string
	Password string

	WindowWidth  int
	WindowHeight int

	Rooms       map[string]*rooms.Room
	CurrentRoom string

	MessageListLayout *listLayouts.QVBoxLayoutWithTriggerSlot
	RoomListLayout    *listLayouts.QRoomVBoxLayoutWithTriggerSlot

	matrixClient
}

type matrixClient struct {
	databases
	Cli    *gomatrix.Client
	Syncer *syncer.MorpheusSyncer
}

// GetCli returns the Matrix Client
func (mc *matrixClient) GetCli() *gomatrix.Client {
	return mc.Cli
}

type databases struct {
	CacheDB db.Storer
}

// SetCurrentRoom sets the new room ID of the MainUI
func (d *databases) SetCacheDB(db db.Storer) {
	d.CacheDB = db
}
