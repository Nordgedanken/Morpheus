package syncer

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/matrix-org/gomatrix"
)

//MorpheusSyncer holds the UserID, the used Storer and the listener
type MorpheusSyncer struct {
	UserID    string
	Store     gomatrix.Storer
	listeners map[string][]OnEventListener // event type to listeners array
	config    *matrix.Config
}

// OnEventListener can be used with DefaultSyncer.OnEventType to be informed of incoming events.
type OnEventListener func(*gomatrix.Event)

// NewMorpheusSyncer returns an instantiated MorpheusSyncer
func NewMorpheusSyncer(userID string, store gomatrix.Storer) *MorpheusSyncer {
	return &MorpheusSyncer{
		UserID:    userID,
		Store:     store,
		listeners: make(map[string][]OnEventListener),
	}
}

// ProcessResponse processes the /sync response in a way suitable for bots. "Suitable for bots" means a stream of
// unrepeating events. Returns a fatal error if a listener panics.
func (s *MorpheusSyncer) ProcessResponse(res *gomatrix.RespSync, since string) (err error) {
	if !s.shouldProcessResponse(res, since) {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ProcessResponse panicked! userID=%s since=%s panic=%s\n%s", s.UserID, since, r, debug.Stack())
		}
	}()

	for roomID, roomData := range res.Rooms.Join {
		room := s.getOrCreateRoom(roomID, "join")
		for _, event := range roomData.State.Events {
			event.RoomID = roomID
			room.UpdateState(&event)
			s.notifyListeners(&event)
		}
		for _, event := range roomData.Timeline.Events {
			event.RoomID = roomID
			s.notifyListeners(&event)
		}
	}
	for roomID, roomData := range res.Rooms.Invite {
		room := s.getOrCreateRoom(roomID, "invite")
		for _, event := range roomData.State.Events {
			event.RoomID = roomID
			room.UpdateState(&event)
			s.notifyListeners(&event)
		}
	}
	for roomID, roomData := range res.Rooms.Leave {
		room := s.getOrCreateRoom(roomID, "leave")
		for _, event := range roomData.Timeline.Events {
			if event.StateKey != nil {
				event.RoomID = roomID
				room.UpdateState(&event)
				s.notifyListeners(&event)
			}
		}
	}
	return
}

// OnEventType allows callers to be notified when there are new events for the given event type.
// There are no duplicate checks.
func (s *MorpheusSyncer) OnEventType(eventType string, callback OnEventListener) {
	_, exists := s.listeners[eventType]
	if !exists {
		s.listeners[eventType] = []OnEventListener{}
	}
	s.listeners[eventType] = append(s.listeners[eventType], callback)
}

// shouldProcessResponse returns true if the response should be processed. May modify the response to remove
// stuff that shouldn't be processed.
func (s *MorpheusSyncer) shouldProcessResponse(resp *gomatrix.RespSync, since string) bool {
	if since == "" {
		return false
	}
	// This is a horrible hack because /sync will return the most recent messages for a room
	// as soon as you /join it. We do NOT want to process those events in that particular room
	// because they may have already been processed (if you toggle the bot in/out of the room).
	//
	// Work around this by inspecting each room's timeline and seeing if an m.room.member event for us
	// exists and is "join" and then discard processing that room entirely if so.
	// TODO: We probably want to process messages from after the last join event in the timeline.
	for roomID, roomData := range resp.Rooms.Join {
		for i := len(roomData.Timeline.Events) - 1; i >= 0; i-- {
			e := roomData.Timeline.Events[i]
			if e.Type == "m.room.member" && e.StateKey != nil && *e.StateKey == s.UserID {
				m := e.Content["membership"]
				mship, ok := m.(string)
				if !ok {
					continue
				}
				if mship == "join" {
					_, ok := resp.Rooms.Join[roomID]
					if !ok {
						continue
					}
					delete(resp.Rooms.Join, roomID)   // don't re-process messages
					delete(resp.Rooms.Invite, roomID) // don't re-process invites
					break
				}
			}
		}
	}
	return true
}

// getOrCreateRoom must only be called by the Sync() goroutine which calls ProcessResponse()
func (s *MorpheusSyncer) getOrCreateRoom(roomID, state string) *gomatrix.Room {
	// Add new Room to the List if new
	_, present := s.config.Rooms[roomID]
	if !present && state == "join" {
		s.config.Rooms[roomID] = matrix.NewRoom(roomID, s.config.GetCli())
		s.config.RoomListLayout.TriggerRoom(roomID)
	}

	room := s.Store.LoadRoom(roomID)
	if room == nil { // create a new Room
		room = gomatrix.NewRoom(roomID)
		s.Store.SaveRoom(room)
	}
	return room
}

func (s *MorpheusSyncer) notifyListeners(event *gomatrix.Event) {
	listeners, exists := s.listeners[event.Type]
	if !exists {
		return
	}
	for _, fn := range listeners {
		fn(event)
	}
}

// OnFailedSync always returns a 10 second wait period between failed /syncs, never a fatal error.
func (s *MorpheusSyncer) OnFailedSync(res *gomatrix.RespSync, err error) (time.Duration, error) {
	return 10 * time.Second, nil
}

// GetFilterJSON returns a filter with a timeline limit of 50.
func (s *MorpheusSyncer) GetFilterJSON(userID string) json.RawMessage {
	return json.RawMessage(`{"room":{"state":{"types":["m.room.*"]},"timeline":{"limit":20,"types":["m.room.message"]}}}`)
}
