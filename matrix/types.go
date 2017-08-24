package matrix

import (
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/gui"
)

// Client is the Client type extender toadd new methods to it which gomatrix lacks currently
type Client struct {
	*gomatrix.Client
}

// RespUserDisplayName is the Response type of getUserDisplayName()
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// GetUserDisplayName returns the Dispaly name to a MXID
func (cli *Client) GetUserDisplayName(mxid string) (displayName string, err error) {
	displayName, err = getUserDisplayName(mxid, cli)
	return
}

// GetOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func (cli *Client) GetOwnUserAvatar() (img *gui.QPixmap) {
	img = getOwnUserAvatar(cli)
	return
}
