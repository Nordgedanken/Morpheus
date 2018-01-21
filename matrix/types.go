package matrix

import (
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/matrix/syncer"
	"github.com/Nordgedanken/Morpheus/ui/listLayouts"
	"github.com/matrix-org/gomatrix"
)

// HTMLMessage is the contents of a Matrix formated message event.
type HTMLMessage struct {
	MsgType       string `json:"msgtype,omitempty"`
	Body          string `json:"body,omitempty"`
	FormattedBody string `json:"formatted_body,omitempty"`
	Format        string `json:"format,omitempty"`
}

// RoomAliases is the json response when getting the room aliases
type RoomAliases struct {
	Age     int `json:"age,omitempty"`
	Content struct {
		Aliases []string `json:"aliases,omitempty"`
	} `json:"content,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	OriginServerTs int64  `json:"origin_server_ts,omitempty"`
	RoomID         string `json:"room_id,omitempty"`
	Sender         string `json:"sender,omitempty"`
	StateKey       string `json:"state_key,omitempty"`
	Type           string `json:"type,omitempty"`
}

// JoinedRooms is the json response when getting the joined rooms list
type JoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms,omitempty"`
}

// RoomAvatar is the json response when getting the room avatar list
type RoomAvatar struct {
	Age     int `json:"age,omitempty"`
	Content struct {
		Info struct {
			H        int    `json:"h,omitempty"`
			Mimetype string `json:"mimetype,omitempty"`
			Size     int    `json:"size,omitempty"`
			W        int    `json:"w,omitempty"`
		} `json:"info,omitempty"`
		URL string `json:"url,omitempty"`
	} `json:"content,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	OriginServerTs int64  `json:"origin_server_ts,omitempty"`
	RoomID         string `json:"room_id,omitempty"`
	Sender         string `json:"sender,omitempty"`
	StateKey       string `json:"state_key,omitempty"`
	Type           string `json:"type,omitempty"`
}

// Config holds important reused information in the UI
type Config struct {
	Username string
	Password string

	WindowWidth  int
	WindowHeight int

	Rooms       map[string]*Room
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
