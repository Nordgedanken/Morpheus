package matrix

import (
	"strings"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
)

//GetClient returns a Client
func GetClient(homeserverURL, userID, accessToken string) (client *gomatrix.Client, err error) {
	userDB, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
	}

	client, ClientErr := gomatrix.NewClient(homeserverURL, userID, accessToken)
	if ClientErr != nil {
		err = ClientErr
		return
	}

	txn := userDB.NewTransaction(true) // Read-write txn
	DBSetAccessTokenErr := txn.Set([]byte("user|accessToken"), []byte(client.AccessToken))
	if DBSetAccessTokenErr != nil {
		err = DBSetAccessTokenErr
		return
	}

	DBSetHomeserverURLErr := txn.Set([]byte("user|homeserverURL"), []byte(client.HomeserverURL.String()))
	if DBSetHomeserverURLErr != nil {
		err = DBSetHomeserverURLErr
		return
	}

	DBSetUserIDErr := txn.Set([]byte("user|userID"), []byte(client.UserID))
	if DBSetUserIDErr != nil {
		err = DBSetUserIDErr
		return
	}

	CommitErr := txn.Commit(nil)
	if CommitErr != nil {
		err = CommitErr
	}
	return
}

//LoginUser Creates a Session for the User
func LoginUser(localpart, password, homeserverURL string) (*gomatrix.Client, error) {
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
		User:     localpart,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	userDB, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)

	txn := userDB.NewTransaction(true) // Read-write txn
	DBSetAccessTokenErr := txn.Set([]byte("user|accessToken"), []byte(resp.AccessToken))
	if DBSetAccessTokenErr != nil {
		return nil, DBSetAccessTokenErr
	}

	DBSetDeviceIDErr := txn.Set([]byte("user|deviceID"), []byte(resp.DeviceID))
	if DBSetDeviceIDErr != nil {
		return nil, DBSetDeviceIDErr
	}

	DBSetHomeserverURLErr := txn.Set([]byte("user|homeserverURL"), []byte(resp.HomeServer))
	if DBSetHomeserverURLErr != nil {
		return nil, DBSetHomeserverURLErr
	}

	DBSetUserIDErr := txn.Set([]byte("user|userID"), []byte(resp.UserID))
	if DBSetUserIDErr != nil {
		return nil, DBSetUserIDErr

	}

	CommitErr := txn.Commit(nil)
	if CommitErr != nil {
		return nil, CommitErr
	}
	return cli, nil
}

// DoLogin generates the needed Client
func DoLogin(localpart, password, homeserverURL, userID, accessToken string, results chan<- *gomatrix.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	var cli *gomatrix.Client
	if accessToken != "" && homeserverURL != "" && userID != "" {
		var cliErr error
		if strings.HasPrefix(homeserverURL, "https://") {
			cli, cliErr = GetClient(homeserverURL, userID, accessToken)
		} else if strings.HasPrefix(homeserverURL, "http://") {
			cli, cliErr = GetClient(homeserverURL, userID, accessToken)
		} else {
			cli, cliErr = GetClient("https://"+homeserverURL, userID, accessToken)
		}
		if cliErr != nil {
			log.Errorln(cliErr)
		}
		cli.SetCredentials(userID, accessToken)
	} else {
		var err error
		cli, err = LoginUser(localpart, password, homeserverURL)
		if err != nil {
			log.Errorln(err)
		}
	}

	results <- cli
}
