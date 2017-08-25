package matrix

import (
	"log"
	"os"
	"strings"

	"github.com/Nordgedanken/Morpheus/util"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/gui"
	"github.com/tidwall/buntdb"
)

var localLog *log.Logger

func init() {
	var file *os.File
	localLog = util.Logger()
	localLog, file = util.StartFileLog(localLog)
	defer file.Close()
}

// getOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func getOwnUserAvatar(cli *gomatrix.Client) *gui.QPixmap {
	// Init local vars
	var avatarData string
	var IMGdata string

	// Get cache
	db.View(func(tx *buntdb.Tx) error {
		QueryErr := tx.AscendKeys("user:"+cli.UserID+":avatarData100x100",
			func(key, value string) bool {
				avatarData = value
				return true
			})
		if QueryErr != nil {
			return QueryErr
		}
		return nil
	})

	//If cache is empty do a ServerQuery
	if avatarData == "" {
		// Get avatarURL
		avatarURL, avatarErr := cli.GetAvatarURL()
		if avatarErr != nil {
			localLog.Println(avatarErr)
		}

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if avatarURL != "" {
			hsURL := cli.HomeserverURL.String()
			avatarURL_splits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURL_splits[0] + "/" + avatarURL_splits[1] + "?width=100&height=100"

			data, err := cli.MakeRequest("GET", urlPath, nil, nil)
			if err != nil {
				localLog.Println(err)
			}
			IMGdata = string(data[:])
		} else {
			//TODO Generate default image (Step: AfterUI)
			IMGdata = "0"
		}

		// Update cache
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			tx.Set("user:"+cli.UserID+":avatarData100x100", IMGdata, nil)
			return nil
		})
		if DBerr != nil {
			localLog.Fatalln(DBerr)
		}
	} else {
		IMGdata = avatarData
	}

	// Convert avatarimage to QPixmap for usage in QT
	avatar := gui.NewQPixmap()
	avatar.LoadFromData(IMGdata, uint(len(IMGdata)), "", 0)
	return avatar
}
