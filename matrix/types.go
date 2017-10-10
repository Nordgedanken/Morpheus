package matrix

// HTMLMessage is the contents of a Matrix formated message event.
type HTMLMessage struct {
	MsgType       string `json:"msgtype"`
	Body          string `json:"body"`
	FormattedBody string `json:"formatted_body"`
	Format        string `json:"format"`
}

type RoomAliases struct {
	Age     int `json:"age"`
	Content struct {
		Aliases []string `json:"aliases"`
	} `json:"content"`
	EventID        string `json:"event_id"`
	OriginServerTs int64  `json:"origin_server_ts"`
	RoomID         string `json:"room_id"`
	Sender         string `json:"sender"`
	StateKey       string `json:"state_key"`
	Type           string `json:"type"`
}

type JoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms"`
}
