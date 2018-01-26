package rooms

import (
	"fmt"
	"strings"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/Nordgedanken/Morpheus/matrix/messages"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/rhinoman/go-commonmark"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
)

const mRoomNameEv = "m.room.name"
const mRoomCanonicalAliasEv = "m.room.canonical_alias"

// Room saves the information of a Room
type Room struct {
	core.QObject

	_                 func(IMGdata []byte) `signal:"SetAvatar"`
	Cli               *gomatrix.Client
	RoomID            string
	RoomName          string
	RoomNameEventType string
	RoomAvatarURL     string
	RoomTopic         string
	Messages          map[string]*messages.Message
}

// GetRooms either returns the joinedRoomsList from the Server or the cachedList
func GetRooms(cli *gomatrix.Client) (rooms []string, err error) {
	// Get Cache
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	var roomsString string
	DBErr := cacheDB.View(func(txn *badger.Txn) error {
		roomsResult, QueryErr := db.Get(txn, []byte("rooms"))
		if QueryErr != nil {
			return QueryErr
		}
		roomsString = fmt.Sprintf("%s", roomsResult)
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	if roomsString == "" {
		roomsResp, ReqErr := cli.JoinedRooms()
		if ReqErr != nil {
			err = ReqErr
			return
		}
		rooms = roomsResp.JoinedRooms
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("rooms"), []byte(strings.Join(rooms, ",")))
			return DBSetErr
		})
		if DBSetErr != nil {
			err = DBSetErr
			return
		}
	} else {
		rooms = strings.Split(roomsString, ",")
	}
	return
}

func (r *Room) AddMessage(message *messages.Message) {
	if r.Messages == nil {
		r.Messages = make(map[string]*messages.Message)
	}
	r.Messages[message.EventID] = message
}

func (r *Room) crawlRoomAvatarURL() {
	roomAvatar := struct {
		URL string `json:"url"`
	}{}
	r.Cli.StateEvent(r.RoomID, "m.room.avatar", "", &roomAvatar)
	r.RoomAvatarURL = roomAvatar.URL
}

func (r *Room) crawlRoomTopic() {
	roomTopic := struct {
		Topic string `json:"topic"`
	}{}
	r.Cli.StateEvent(r.RoomID, "m.room.topic", "", &roomTopic)
	r.RoomTopic = roomTopic.Topic
}

// GetRoomTopic returns the Topic of the Room and crawls it if needed
func (r *Room) GetRoomTopic() (topic string) {
	if r.RoomTopic == "" {
		r.crawlRoomTopic()
	}

	mardownMessage := commonmark.Md2Html(r.RoomTopic, 0)

	topic = r.RoomTopic
	if mardownMessage == r.RoomTopic {
		topic = r.RoomTopic
	} else {
		r.RoomTopic = mardownMessage
	}
	return
}

// GetRoomAvatar generates the Avatar Image for a Room
func (r *Room) GetRoomAvatar() {
	// Get the Avatar URL if needed
	if r.RoomAvatarURL == "" {
		r.crawlRoomAvatarURL()
	}

	// Get the image Data
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	// Init local vars
	var roomAvatarData []byte
	var IMGdata []byte

	// Get cache
	DBErr := cacheDB.View(func(txn *badger.Txn) error {
		roomAvatarDataResult, QueryErr := db.Get(txn, []byte("room|"+r.RoomID+"|84x84"))
		if QueryErr != nil {
			return QueryErr
		}
		roomAvatarData = roomAvatarDataResult
		return nil
	})
	if DBErr != nil {
		log.Errorf("%s", DBErr)
		return
	}

	//If cache is empty do a ServerQuery
	if len(roomAvatarData) <= 0 {
		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 84x84. Else make the data string empty
		if r.RoomAvatarURL != "" {
			hsURL := r.Cli.HomeserverURL.String()
			roomAvatarURLSplits := strings.Split(strings.Replace(r.RoomAvatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + roomAvatarURLSplits[0] + "/" + roomAvatarURLSplits[1] + "?width=84&height=84&method=crop"

			data, ReqErr := r.Cli.MakeRequest("GET", urlPath, nil, nil)
			if ReqErr != nil {
				log.Errorf("%s", ReqErr)
				return
			}
			IMGdata = data
		} else {
			log.Println("Generating Room Avatar")
			var GenerateImgErr error
			var roomName string
			if r.RoomName == "" {
				r.crawlRoomName()
			}
			roomName = r.RoomName
			log.Println(roomName)
			if roomName == "" {
				roomName = "#"
			}
			IMGdata, GenerateImgErr = matrix.GenerateGenericImages(roomName, 84)
			if GenerateImgErr != nil {
				log.Errorf("%s", GenerateImgErr)
				return
			}
		}

		// Update cache
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("room|"+r.RoomID+"|84x84"), IMGdata)
			return DBSetErr
		})
		if DBSetErr != nil {
			log.Errorf("%s", DBSetErr)
			return
		}

	} else {
		IMGdata = roomAvatarData
	}

	r.SetAvatar(IMGdata)
	return
}

func (r *Room) crawlRoomName() {
	roomName := struct {
		Name string `json:"name"`
	}{}
	roomCanonicalAlias := struct {
		Alias string `json:"alias"`
	}{}

	if roomNameStateEventErr := r.Cli.StateEvent(r.RoomID, mRoomNameEv, "", &roomName); roomNameStateEventErr != nil {
		log.Println(roomNameStateEventErr)
		// Not returning as a Error NotFound is allowed
	}
	if roomName.Name == "" {
		if roomCanonicalAliasStateEventErr := r.Cli.StateEvent(r.RoomID, mRoomCanonicalAliasEv, "", &roomCanonicalAlias); roomCanonicalAliasStateEventErr != nil {
			log.Println(roomCanonicalAliasStateEventErr)
			// Not returning as a Error NotFound is allowed
		}
		if roomCanonicalAlias.Alias == "" {
			r.RoomNameEventType = "roomID"
			r.RoomName = r.RoomID
		} else {
			r.RoomNameEventType = mRoomCanonicalAliasEv
			r.RoomName = roomCanonicalAlias.Alias
		}
	} else {
		r.RoomNameEventType = mRoomNameEv
		r.RoomName = roomName.Name
	}

}

// UpdateRoomNameByEvent used to Update the Room Name of a Room when a Room Change Event comes down the Sync
func (r *Room) UpdateRoomNameByEvent(newName, evType string) {
	if r.RoomNameEventType == "" {
		r.getRoomNameEventTypeFromDB()
	}
	if r.RoomNameEventType == mRoomNameEv && r.RoomNameEventType == evType {
		r.RoomName = newName
		r.cacheRoomName()
	} else if r.RoomNameEventType != mRoomNameEv {
		r.RoomName = newName
		r.cacheRoomName()
	}
}

func (r *Room) cacheRoomName() (err error) {
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}

	// Update cache
	DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
		DBSetErr := txn.Set([]byte("room|"+r.RoomID+"|nameEventType"), []byte(r.RoomNameEventType))
		if DBSetErr != nil {
			return DBSetErr
		}

		DBSetErr = txn.Set([]byte("room|"+r.RoomID+"|name"), []byte(r.RoomName))
		return DBSetErr
	})
	if DBSetErr != nil {
		err = DBSetErr
		return
	}
	return
}

func (r *Room) getRoomNameEventTypeFromDB() (err error) {
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}

	cacheDB.View(func(txn *badger.Txn) error {
		roomNameEventTypeResult, QueryErr := db.Get(txn, []byte("room|"+r.RoomID+"|nameEventType"))
		if QueryErr != nil {
			return QueryErr
		}
		r.RoomNameEventType = fmt.Sprintf("%s", roomNameEventTypeResult)
		return nil
	})

	return
}

func (r *Room) getRoomNameFromDB() (err error) {
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}

	cacheDB.View(func(txn *badger.Txn) error {
		roomNameResult, QueryErr := db.Get(txn, []byte("room|"+r.RoomID+"|name"))
		if QueryErr != nil {
			return QueryErr
		}
		r.RoomName = fmt.Sprintf("%s", roomNameResult)
		return nil
	})

	return
}

// GetRoomName gives you the name of the current Room
func (r *Room) GetRoomName() (name string) {
	r.getRoomNameFromDB()
	if r.RoomName == "" {
		r.crawlRoomName()
		r.cacheRoomName()
	}
	name = r.RoomName
	return
}
