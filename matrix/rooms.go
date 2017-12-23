package matrix

import (
	"bytes"
	"image"
	"image/png"
	"strings"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/rhinoman/go-commonmark"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/gui"
)

// Room saves the information of a Room
type Room struct {
	cli           *gomatrix.Client
	RoomID        string
	RoomName      string
	RoomAvatarURL string
	RoomTopic     string
}

// NewRoom Inits a new Room struct
func NewRoom(roomID string, cli *gomatrix.Client) (room *Room) {
	room = &Room{RoomID: roomID, cli: cli}
	return
}

func (r *Room) crawlRoomAvatarURL() {
	roomAvatar := struct {
		URL string `json:"url"`
	}{}
	r.cli.StateEvent(r.RoomID, "m.room.avatar", "", &roomAvatar)
	r.RoomAvatarURL = roomAvatar.URL
}

func (r *Room) crawlRoomTopic() {
	roomTopic := struct {
		Topic string `json:"topic"`
	}{}
	r.cli.StateEvent(r.RoomID, "m.room.topic", "", &roomTopic)
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
func (r *Room) GetRoomAvatar() (avatarResp *gui.QPixmap, err error) {
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
		roomAvatarDataItem, QueryErr := txn.Get([]byte("room|" + r.RoomID + "|84x84"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			roomAvatarDataBytes, roomAvatarDataErr := roomAvatarDataItem.Value()
			roomAvatarData = roomAvatarDataBytes
			return roomAvatarDataErr
		}
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	//If cache is empty do a ServerQuery
	if len(roomAvatarData) <= 0 {
		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if r.RoomAvatarURL != "" {
			hsURL := r.cli.HomeserverURL.String()
			roomAvatarURLSplits := strings.Split(strings.Replace(r.RoomAvatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + roomAvatarURLSplits[0] + "/" + roomAvatarURLSplits[1] + "?width=84&height=84&method=crop"

			data, ReqErr := r.cli.MakeRequest("GET", urlPath, nil, nil)
			if ReqErr != nil {
				err = ReqErr
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
			IMGdata, GenerateImgErr = generateGenericImages(roomName, 84)
			if GenerateImgErr != nil {
				err = GenerateImgErr
				return
			}
		}

		// Update cache
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("room|"+r.RoomID+"|84x84"), IMGdata)
			return DBSetErr
		})
		if DBSetErr != nil {
			err = DBSetErr
			return
		}

	} else {
		IMGdata = roomAvatarData
	}

	reader := bytes.NewReader(IMGdata)
	srcIMG, _, DecodeErr := image.Decode(reader)
	if DecodeErr != nil {
		err = DecodeErr
	}

	avatar := gui.NewQPixmap()
	buf := new(bytes.Buffer)
	ConvErr := png.Encode(buf, srcIMG)
	if ConvErr != nil {
		err = ConvErr
	}

	str := buf.String()
	avatar.LoadFromData(str, uint(len(str)), "", 0)
	avatarResp = avatar
	return
}

func (r *Room) crawlRoomName() {
	roomName := struct {
		Name string `json:"name"`
	}{}
	roomCanoncialAlias := struct {
		Alias string `json:"alias"`
	}{}

	if roomNameStateEventErr := r.cli.StateEvent(r.RoomID, "m.room.name", "", &roomName); roomNameStateEventErr != nil {
		log.Println(roomNameStateEventErr)
		// Not returning as a Error NotFound is allowed
	}
	if roomName.Name == "" {
		if roomCanoncialAliasStateEventErr := r.cli.StateEvent(r.RoomID, "m.room.canonical_alias", "", &roomCanoncialAlias); roomCanoncialAliasStateEventErr != nil {
			log.Println(roomCanoncialAliasStateEventErr)
			// Not returning as a Error NotFound is allowed
		}
		if roomCanoncialAlias.Alias == "" {
			r.RoomName = r.RoomID
		} else {
			r.RoomName = roomCanoncialAlias.Alias
		}
	} else {
		r.RoomName = roomName.Name
	}

}

// GetRoomName gives you the name of the current Room
func (r *Room) GetRoomName() (name string) {
	if r.RoomName == "" {
		r.crawlRoomName()
	}
	name = r.RoomName
	return
}
