package matrix

import (
	"log"
	"strings"
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/matrix-org/gomatrix"
	"github.com/tidwall/buntdb"
)

//GetClient returns a Client
func GetClient(homeserverURL, userID, accessToken string) (client *gomatrix.Client, err error) {
	db, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}
	defer db.Close()

	client, ClientErr := gomatrix.NewClient(homeserverURL, userID, accessToken)
	if ClientErr != nil {
		err = ClientErr
		return
	}

	DBErr := db.Update(func(tx *buntdb.Tx) error {
		_, _, DBSetAccessTokenErr := tx.Set("user|accessToken", client.AccessToken, nil)
		if DBSetAccessTokenErr != nil {
			return DBSetAccessTokenErr
		}

		_, _, DBSetHomeserverURLErr := tx.Set("user|homeserverURL", client.HomeserverURL.String(), nil)
		if DBSetHomeserverURLErr != nil {
			return DBSetHomeserverURLErr
		}

		_, _, DBSetUserIDErr := tx.Set("user|userID", client.UserID, nil)
		return DBSetUserIDErr
	})
	if DBErr != nil {
		err = DBErr
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

	db, DBOpenErr := db.OpenUserDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}
	defer db.Close()
	cli.SetCredentials(resp.UserID, resp.AccessToken)
	DBerr := db.Update(func(tx *buntdb.Tx) error {
		_, _, DBSetAccessTokenErr := tx.Set("user|accessToken", resp.AccessToken, nil)
		if DBSetAccessTokenErr != nil {
			return DBSetAccessTokenErr
		}

		_, _, DBSetDeviceIDErr := tx.Set("user|deviceID", resp.DeviceID, nil)
		if DBSetDeviceIDErr != nil {
			return DBSetDeviceIDErr
		}

		_, _, DBSetHomeserverURLErr := tx.Set("user|homeserverURL", resp.HomeServer, nil)
		if DBSetHomeserverURLErr != nil {
			return DBSetHomeserverURLErr
		}

		_, _, DBSetUserIDErr := tx.Set("user|userID", resp.UserID, nil)
		return DBSetUserIDErr
	})
	if DBerr != nil {
		return nil, DBerr
	}
	return cli, nil
}

// DoLogin generates the needed Client
func DoLogin(username, password, homeserverURL, userID, accessToken string, localLog *log.Logger, results chan<- *gomatrix.Client, wg *sync.WaitGroup) {
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
			localLog.Println(cliErr)
		}
		cli.SetCredentials(userID, accessToken)
	} else {
		var err error
		cli, err = LoginUser(username, password)
		if err != nil {
			localLog.Println(err)
		}
	}

	results <- cli
}
