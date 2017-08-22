package matrix

import (
	"github.com/matrix-org/gomatrix"
)

// RespUserDisplayName is the Response type of GetUserDisplayName()
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// RespOwnUserAvatar is the Response type of GetUserDisplayName()
// type RespOwnUserAvatar struct {
// 	DisplayName string `json:"displayname"`
// }

// GetUserDisplayName returns the Dispaly name to a MXID
func GetUserDisplayName(mxid string, cli *gomatrix.Client) (resp *RespUserDisplayName, err error) {
	urlPath := cli.BuildURL("profile", mxid, "displayname")
	_, err = cli.MakeRequest("GET", urlPath, nil, &resp)
	return
}

// GetOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
// func GetOwnUserAvatar(cli *gomatrix.Client) *gui.QPixmap {
// 	avatarURL, avatarErr := cli.GetAvatarURL()
// 	if avatarErr != nil {
// 		localLog.Println(avatarErr)
// 	}
//
// 	avatar := gui.NewQPixmap()
// 	avatar.LoadFromData()
// }
