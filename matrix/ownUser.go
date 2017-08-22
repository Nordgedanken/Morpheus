package matrix

import (
	"log"
	"os"
	"strings"

	"github.com/Nordgedanken/Neo/util"
	"github.com/therecipe/qt/gui"
	"github.com/tidwall/buntdb"
)

var localLog *log.Logger

// getUserDisplayName returns the Dispaly name to a MXID
func getUserDisplayName(mxid string, cli *Client) (resp *RespUserDisplayName, err error) {

	urlPath := cli.BuildURL("profile", mxid, "displayname")
	_, err = cli.MakeRequest("GET", urlPath, nil, &resp)
	err = db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:displayName", resp.DisplayName, nil)
		return nil
	})
	return
}

// getOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func getOwnUserAvatar(cli *Client) *gui.QPixmap {
	var file *os.File
	localLog = util.Logger()
	localLog, file = util.StartFileLog(localLog)
	defer file.Close()
	avatarURL, avatarErr := cli.GetAvatarURL()
	if avatarErr != nil {
		localLog.Println(avatarErr)
	}

	var IMGdata string
	if avatarURL != "" {
		hsURL := cli.HomeserverURL.String()
		avatarURL_splits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

		urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURL_splits[0] + "/" + avatarURL_splits[1] + "?width=100&height=100"
		println(urlPath)

		data, err := cli.MakeRequest("GET", urlPath, nil, nil)
		if err != nil {
			localLog.Println(err)
		}
		IMGdata = string(data[:])
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			tx.Set("user:avatarData100x100", IMGdata, nil)
			return nil
		})
		if DBerr != nil {
			localLog.Fatalln(err)
		}
	} else {
		IMGdata = ""
	}

	avatar := gui.NewQPixmap()
	avatar.LoadFromData(IMGdata, uint(len(IMGdata)), "", 0)
	return avatar
}
