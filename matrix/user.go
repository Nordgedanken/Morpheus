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

// GenerateGenericImages generates a byte slice containing a Image with the starting char as symbol and a colored background
func GenerateGenericImages(identifier string, size int) (imgData []byte, err error) {
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
		avatarDataResult, QueryErr := db.Get(txn, []byte("user|"+mxid+"|avatarData"+strconv.Itoa(size)+"x"+strconv.Itoa(size)))
		if QueryErr != nil {
			return QueryErr
		}
		avatarData = avatarDataResult
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	//If cache is empty do a ServerQuery
	log.Println(len(avatarData))
	if len(avatarData) <= 0 {
		log.Println("Download")
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
			IMGdata, GenerateImgErr = GenerateGenericImages(DisplayName, size)
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
		accessTokenResult, QueryErr := db.Get(txn, []byte("user|accessToken"))
		if QueryErr != nil {
			return QueryErr
		}
		accessToken = fmt.Sprintf("%s", accessTokenResult)

		homeserverURLResult, QueryErr := db.Get(txn, []byte("user|homeserverURL"))
		if QueryErr != nil {
			return QueryErr
		}
		homeserverURL = fmt.Sprintf("%s", homeserverURLResult)

		userIDResult, QueryErr := db.Get(txn, []byte("user|userID"))
		if QueryErr != nil {
			return QueryErr
		}
		userID = fmt.Sprintf("%s", userIDResult)
		return nil
	})
	if DBErr != nil {
		err = DBErr
	}

	return
}
