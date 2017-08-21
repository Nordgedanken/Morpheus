package matrix

import (
	"strings"
	"sync"

	"github.com/matrix-org/gomatrix"
)

var clientInstance *gomatrix.Client
var once sync.Once

//getClient returns a Client
func getClient(homeserverURL, userID, accessToken string) (*gomatrix.Client, error) {
	var err error
	once.Do(func() {
		clientInstance, err = gomatrix.NewClient(homeserverURL, userID, accessToken)
	})
	if err != nil {
		return nil, err
	}
	return clientInstance, nil
}

//LoginUser Creates a Session for the User
func LoginUser(username, password string) (*gomatrix.Client, error) {
	usernameSplit := strings.Split(username, ":")
	homeserverURL := usernameSplit[1]
	cli, cliErr := getClient(homeserverURL, "", "")
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
