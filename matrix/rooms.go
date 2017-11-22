package matrix

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"strings"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
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
	topic = r.RoomTopic
	return
}

// GetRoomAvatar generates the Avatar Image for a Room
func (r *Room) GetRoomAvatar() (avatarResp *gui.QPixmap, err error) {
	// Get the Avatar URL if needed
	if r.RoomAvatarURL == "" {
		r.crawlRoomAvatarURL()
	}

	// Get the image Data
	db, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}

	// Init local vars
	var roomAvatarData string
	var IMGdata string

	// Get cache
	DBErr := db.View(func(txn *badger.Txn) error {
		roomAvatarDataItem, QueryErr := txn.Get([]byte("room|" + r.RoomID + "|84x84"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			roomAvatarDataBytes, roomAvatarDataErr := roomAvatarDataItem.Value()
			roomAvatarData = fmt.Sprintf("%s", roomAvatarDataBytes)
			return roomAvatarDataErr
		}
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	//If cache is empty do a ServerQuery
	if roomAvatarData == "" {
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
			IMGdata = string(data[:])
		} else {
			localLog.Println("Generating Room Avatar")
			var GenerateImgErr error
			var roomName string
			if r.RoomName == "" {
				r.crawlRoomName()
			}
			roomName = r.RoomName
			localLog.Println(roomName)
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
		txn := db.NewTransaction(true) // Read-write txn
		DBSetErr := txn.Set([]byte("room|"+r.RoomID+"|84x84"), []byte(IMGdata))
		if DBSetErr != nil {
			err = DBSetErr
			return
		}
	} else {
		IMGdata = roomAvatarData
	}

	reader := bytes.NewReader([]byte(IMGdata))
	srcIMG, _, DecodeErr := image.Decode(reader)
	if DecodeErr != nil {
		err = DecodeErr
	}

	// Convert avatarimage to QPixmap for usage in QT
	canvas := image.NewRGBA(srcIMG.Bounds())
	cx := srcIMG.Bounds().Min.X + srcIMG.Bounds().Dx()/2
	cy := srcIMG.Bounds().Min.Y + srcIMG.Bounds().Dy()/2
	draw.DrawMask(canvas, canvas.Bounds(), srcIMG, image.ZP, &circle{image.Point{cx, cy}, cx}, image.ZP, draw.Over)

	avatar := gui.NewQPixmap()
	buf := new(bytes.Buffer)
	ConvErr := png.Encode(buf, canvas)
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
		localLog.Println(roomNameStateEventErr)
		// Not returning as a Error NotFound is allowed
	}
	if roomName.Name == "" {
		if roomCanoncialAliasStateEventErr := r.cli.StateEvent(r.RoomID, "m.room.canonical_alias", "", &roomCanoncialAlias); roomCanoncialAliasStateEventErr != nil {
			localLog.Println(roomCanoncialAliasStateEventErr)
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
