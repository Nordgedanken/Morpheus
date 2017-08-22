package matrix

import (
	"log"
	"os"
	"strings"

	"github.com/Nordgedanken/Neo/util"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/gui"
)

var localLog *log.Logger

// RespUserDisplayName is the Response type of GetUserDisplayName()
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// GetUserDisplayName returns the Dispaly name to a MXID
func GetUserDisplayName(mxid string, cli *gomatrix.Client) (resp *RespUserDisplayName, err error) {
	urlPath := cli.BuildURL("profile", mxid, "displayname")
	_, err = cli.MakeRequest("GET", urlPath, nil, &resp)
	return
}

// GetOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetOwnUserAvatar(cli *gomatrix.Client) *gui.QPixmap {
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
	} else {
		IMGdata = ""
	}

	avatar := gui.NewQPixmap()
	avatar.LoadFromData(IMGdata, uint(len(IMGdata)), "", 0)
	return avatar
}
