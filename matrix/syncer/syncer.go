package syncer

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/Nordgedanken/Morpheus/matrix/globalTypes"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
)

//MorpheusSyncer holds the UserID, the used Storer and the listener
type MorpheusSyncer struct {
	UserID    string
	Store     gomatrix.Storer
	listeners map[string][]OnEventListener // event type to listeners array
	config    *globalTypes.Config
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
	log.Infoln("Since: ", since)
	log.Infof("Res: %+v\n", res)
	if !s.shouldProcessResponse(res, since) {
		log.Infoln("bug")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ProcessResponse panicked! userID=%s since=%s panic=%s\n%s", s.UserID, since, r, debug.Stack())
		}
	}()

	for roomID, roomData := range res.Rooms.Join {
		log.Infoln("Join Loop")
		log.Infof("roomID: %+v\n", roomID)
		log.Infof("roomData: %+v\n", roomData)
		room := s.getOrCreateRoom(roomID)
		log.Infof("Room: %+v\n", room)
		for _, event := range roomData.State.Events {
			log.Infoln("Join Event")
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
		log.Infoln("Invite Loop")
		room := s.getOrCreateRoom(roomID)
		for _, event := range roomData.State.Events {
			log.Infoln("Invite Event")
			event.RoomID = roomID
			room.UpdateState(&event)
			s.notifyListeners(&event)
		}
	}
	for roomID, roomData := range res.Rooms.Leave {
		log.Infoln("Leave Loop")
		room := s.getOrCreateRoom(roomID)
		for _, event := range roomData.Timeline.Events {
			if event.StateKey != nil {
				log.Infoln("Leave Event")
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
func (s *MorpheusSyncer) getOrCreateRoom(roomID string) *gomatrix.Room {
	// Add new Room to the List if new
	log.Infoln(s.config.Rooms)

	room := s.config.Rooms[roomID]
	gomatrixRoom := gomatrix.NewRoom(roomID)
	if room == nil { // create a new Room
		s.config.RoomList.RoomCount++
		if (s.config.RoomList.RoomCount % 10) == 0 {
			s.config.App.ProcessEvents(core.QEventLoop__AllEvents)
		}
		s.config.Rooms[roomID] = room
		go s.config.RoomList.TriggerRoom(roomID)
	}
	return gomatrixRoom
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
	log.Errorln(err)
	return 10 * time.Second, nil
}

// GetFilterJSON returns a filter with a timeline limit of 50.
func (s *MorpheusSyncer) GetFilterJSON(userID string) json.RawMessage {
	//return json.RawMessage(`{"room":{"timeline":{"limit":50}}}`)
	return json.RawMessage(`{"room":{"state":{"types":["m.room.*"]},"timeline":{"limit":20,"types":["m.room.message"]}}}`)
}
