package messages

import (
	"strings"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	"github.com/opennota/linkify"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
)

// Message saves the information of a Message
type Message struct {
	core.QObject
	setAvatarFuncs []func(IMGdata []byte)
	Cli            *gomatrix.Client
	EventID        string
	Author         string
	EventType      string
	AvatarURL      string
	Message        string
	Timestamp      int64
}

func (m *Message) crawlAvatarURL() (err error) {
	// Get avatarURL
	urlPath := m.Cli.BuildURL("profile", m.Author, "avatar_url")
	s := struct {
		AvatarURL string `json:"avatar_url"`
	}{}

	_, ReqErr := m.Cli.MakeRequest("GET", urlPath, nil, &s)
	if ReqErr != nil {
		err = ReqErr
		return
	}
	m.AvatarURL = s.AvatarURL
	return
}

// Linkify makes urls to html links
func (m *Message) Linkify() (err error) {
	lm := linkify.Links(m.Message)
	for _, l := range lm {
		link := m.Message[l.Start:l.End]
		if l.Start-10 > 0 {
			log.Println(m.Message[l.Start-10 : l.Start])
			log.Println(m.Message[l.Start-(1+l.Start+l.End) : l.Start])
			if !strings.Contains(m.Message[l.Start-10:l.Start], "<a href='") {
				if l.Start-(1+l.Start+l.End) > 0 {
					if !strings.Contains(m.Message[l.Start-(1+l.Start+l.End):l.Start], "<a href='"+link+"'>") {
						m.Message = strings.Replace(m.Message, link, "<a href='"+link+"'>"+link+"</a>", -1)
					}
				} else if l.Start-(1+l.Start+l.End) <= 0 {
					if !strings.Contains(m.Message[0:l.Start], "<a href='"+link+"'>") {
						m.Message = strings.Replace(m.Message, link, "<a href='"+link+"'>"+link+"</a>", -1)
					}
				}
			}
		} else if l.Start-10 <= 0 {
			if !strings.Contains(m.Message[0:l.Start], "<a href='") {
				if !strings.Contains(m.Message[0:l.Start], "<a href='"+link+"'>") {
					m.Message = strings.Replace(m.Message, link, "<a href='"+link+"'>"+link+"</a>", -1)
				}
			}
		}
	}
	return
}

// GetUserAvatar returns a *gui.QPixmap of an UserAvatar
func (m *Message) GetUserAvatar() {
	// Get the Avatar URL if needed
	if m.AvatarURL == "" {
		m.crawlAvatarURL()
	}

	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	// Init local vars
	var avatarData []byte
	var IMGdata []byte

	// Get cache
	DBErr := cacheDB.View(func(txn *badger.Txn) error {
		avatarDataResult, QueryErr := db.Get(txn, []byte("user|"+m.Author+"|avatarData61x61"))
		if QueryErr != nil {
			return QueryErr
		}
		avatarData = avatarDataResult
		return nil
	})
	if DBErr != nil {
		log.Errorf("%s", DBErr)
		return
	}

	//If cache is empty do a ServerQuery
	if len(avatarData) <= 0 {

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if m.AvatarURL != "" {
			hsURL := m.Cli.HomeserverURL.String()
			avatarURLSplits := strings.Split(strings.Replace(m.AvatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURLSplits[0] + "/" + avatarURLSplits[1] + "?width=61&height=61&method=crop"

			data, ReqErr := m.Cli.MakeRequest("GET", urlPath, nil, nil)
			if ReqErr != nil {
				log.Errorf("%s", ReqErr)
				return
			}
			IMGdata = data
		} else {
			DisplayNameResp, _ := m.Cli.GetDisplayName(m.Author)
			DisplayName := DisplayNameResp.DisplayName
			var GenerateImgErr error
			IMGdata, GenerateImgErr = matrix.GenerateGenericImages(DisplayName, 61)
			if GenerateImgErr != nil {
				log.Errorf("%s", GenerateImgErr)
				return
			}
		}

		// Update cache
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("user|"+m.Author+"|avatarData61x61"), IMGdata)
			return DBSetErr
		})
		if DBSetErr != nil {
			log.Errorf("%s", DBSetErr)
			return
		}
	} else {
		IMGdata = avatarData
	}

	m.SetAvatar(IMGdata)
	return
}

// ConnectSetAvatar registers a callback function
func (m *Message) ConnectSetAvatar(f func(IMGdata []byte)) {
	m.setAvatarFuncs = append(m.setAvatarFuncs, f)
	return
}

// SetAvatar triggers all callback functions
func (m *Message) SetAvatar(IMGdata []byte) {
	for _, f := range m.setAvatarFuncs {
		f(IMGdata)
	}
	return
}
