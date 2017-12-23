package matrix

import (
	"bytes"
	"fmt"
	"image/png"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/disintegration/letteravatar"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/gui"
)

func generateGenericImages(identifier string, size int) (imgData []byte, err error) {
	if (identifier[0] == '#' || identifier[0] == '!' || identifier[0] == '@') && len(identifier) > 1 {
		identifier = identifier[1:]
	}

	avatarChar, _ := utf8.DecodeRuneInString(identifier)
	img, LetterAvatarErr := letteravatar.Draw(size, unicode.ToUpper(avatarChar), nil)

	if LetterAvatarErr != nil {
		err = LetterAvatarErr
		return
	}

	buf := new(bytes.Buffer)
	EncodeErr := png.Encode(buf, img)
	if EncodeErr != nil {
		err = EncodeErr
		return
	}
	imgData = buf.Bytes()
	return
}

// GetOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetOwnUserAvatar(cli *gomatrix.Client) (avatar *gui.QPixmap, err error) {
	avatar, err = GetUserAvatar(cli, cli.UserID, 61)
	return
}

// GetUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetUserAvatar(cli *gomatrix.Client, mxid string, size int) (avatarResp *gui.QPixmap, err error) {
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	// Init local vars
	var avatarData []byte
	var IMGdata []byte

	// Get cache
	DBErr := cacheDB.View(func(txn *badger.Txn) error {
		roomAvatarDataItem, QueryErr := txn.Get([]byte("user|" + mxid + "|avatarData" + strconv.Itoa(size) + "x" + strconv.Itoa(size)))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			avatarDataBytes, avatarDataErr := roomAvatarDataItem.Value()
			avatarData = avatarDataBytes
			return avatarDataErr
		}
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	//If cache is empty do a ServerQuery
	if len(avatarData) <= 0 {
		// Get avatarURL
		urlPath := cli.BuildURL("profile", mxid, "avatar_url")
		s := struct {
			AvatarURL string `json:"avatar_url"`
		}{}

		_, ReqErr := cli.MakeRequest("GET", urlPath, nil, &s)
		if ReqErr != nil {
			err = ReqErr
			return
		}
		avatarURL := s.AvatarURL

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if avatarURL != "" {
			hsURL := cli.HomeserverURL.String()
			avatarURLSplits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURLSplits[0] + "/" + avatarURLSplits[1] + "?width=" + strconv.Itoa(size) + "&height=" + strconv.Itoa(size) + "&method=crop"

			data, ReqErr := cli.MakeRequest("GET", urlPath, nil, nil)
			if ReqErr != nil {
				err = ReqErr
				return
			}
			IMGdata = data
		} else {
			DisplayNameResp, _ := cli.GetDisplayName(mxid)
			DisplayName := DisplayNameResp.DisplayName
			var GenerateImgErr error
			IMGdata, GenerateImgErr = generateGenericImages(DisplayName, size)
			if GenerateImgErr != nil {
				err = GenerateImgErr
				return
			}
		}

		// Update cache
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("user|"+mxid+"|avatarData"+strconv.Itoa(size)+"x"+strconv.Itoa(size)), IMGdata)
			return DBSetErr
		})
		if DBSetErr != nil {
			err = DBSetErr
			return
		}
	} else {
		IMGdata = avatarData
	}

	avatar := gui.NewQPixmap()

	str := string(IMGdata[:])

	avatar.LoadFromData(string(str[:]), uint(len(str)), "", 0)
	avatarResp = avatar
	return
}

//GetUserDataFromCache gets the last User from the cache
func GetUserDataFromCache() (accessToken, homeserverURL, userID string, err error) {
	UserDB, err := db.OpenUserDB()

	// Get cache
	DBErr := UserDB.View(func(txn *badger.Txn) error {
		accessTokenItem, QueryErr := txn.Get([]byte("user|accessToken"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			accessTokenByte, accessTokenErr := accessTokenItem.Value()
			accessToken = fmt.Sprintf("%s", accessTokenByte)
			if accessTokenErr != nil {
				return accessTokenErr
			}
		}

		homeserverURLItem, QueryErr := txn.Get([]byte("user|homeserverURL"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			homeserverURLByte, homeserverURLErr := homeserverURLItem.Value()
			homeserverURL = fmt.Sprintf("%s", homeserverURLByte)
			if homeserverURLErr != nil {
				return homeserverURLErr
			}
		}

		userIDItem, QueryErr := txn.Get([]byte("user|userID"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			userIDByte, userIDErr := userIDItem.Value()
			userID = fmt.Sprintf("%s", userIDByte)
			return userIDErr
		}
		return nil
	})
	if DBErr != nil {
		err = DBErr
	}

	return
}
