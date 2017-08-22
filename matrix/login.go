package matrix

import (
	"strings"
	"sync"

	"github.com/matrix-org/gomatrix"
	"github.com/tidwall/buntdb"
)

var clientInstance *Client
var once sync.Once

//getClient returns a Client
func getClient(homeserverURL, userID, accessToken string) (*Client, error) {
	var err error
	var client *gomatrix.Client
	once.Do(func() {
		client, err = gomatrix.NewClient(homeserverURL, userID, accessToken)
	})
	if err != nil {
		return nil, err
	}
	clientInstance = &Client{client}

	DBerr := db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:accessToken", clientInstance.AccessToken, nil)
		tx.Set("user:homeserverURL", clientInstance.HomeserverURL.String(), nil)
		tx.Set("user:userID", clientInstance.UserID, nil)
		return nil
	})
	if DBerr != nil {
		return nil, DBerr
	}

	return clientInstance, nil
}

//LoginUser Creates a Session for the User
func LoginUser(username, password string) (*Client, error) {
	usernameSplit := strings.Split(username, ":")
	homeserverURL := usernameSplit[1]
	cli, cliErr := getClient("https://"+homeserverURL, "", "")
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
	return cli, nil
}
