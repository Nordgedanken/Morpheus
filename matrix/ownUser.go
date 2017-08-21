package matrix

import (
	"github.com/matrix-org/gomatrix"
)

type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

func GetUserDisplayName(mxid string, cli *gomatrix.Client) (resp *RespUserDisplayName, err error) {
	urlPath := cli.BuildURL("profile", mxid, "displayname")
	_, err = cli.MakeRequest("GET", urlPath, nil, &resp)
	return
}
