package matrix

import (
	"strings"

	"github.com/matrix-org/gomatrix"
	"github.com/tidwall/buntdb"
)

//GetClient returns a Client
func GetClient(homeserverURL, userID, accessToken string) (client *gomatrix.Client, err error) {
	client, err = gomatrix.NewClient(homeserverURL, userID, accessToken)
	if err != nil {
		client = nil
		return
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:accessToken", client.AccessToken, nil)
		tx.Set("user:homeserverURL", client.HomeserverURL.String(), nil)
		tx.Set("user:userID", client.UserID, nil)
		return nil
	})
	if err != nil {
		client = nil
		return
	}

	return
}

//LoginUser Creates a Session for the User
func LoginUser(username, password string) (*gomatrix.Client, error) {
	usernameSplit := strings.Split(username, ":")
	homeserverURL := usernameSplit[1]
	var cli *gomatrix.Client
	var cliErr error
	if strings.HasPrefix(homeserverURL, "https://") {
		cli, cliErr = GetClient(homeserverURL, "", "")
	} else if strings.HasPrefix(homeserverURL, "http://") {
		cli, cliErr = GetClient(homeserverURL, "", "")
	} else {
		cli, cliErr = GetClient("https://"+homeserverURL, "", "")
	}
	if cliErr != nil {
		return nil, cliErr
	}

	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)
	DBerr := db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:accessToken", resp.AccessToken, nil)
		tx.Set("user:deviceID", resp.DeviceID, nil)
		tx.Set("user:homeserverURL", resp.HomeServer, nil)
		tx.Set("user:userID", resp.UserID, nil)
		return nil
	})
	if DBerr != nil {
		return nil, DBerr
	}
	return cli, nil
}
