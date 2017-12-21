package scalar

import (
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
)

type openIDTokenResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	MatrixServerName string `json:"matrix_server_name"`
}

func getOpenIDToken(cli *gomatrix.Client) (resp *openIDTokenResponse, err error) {
	urlPath := cli.BuildURL("user", cli.UserID, "openid", "request_token")
	_, err = makeRequest(cli, "POST", urlPath, "{}", &resp)
	return
}

type scalarRegisterResp struct {
	ScalarToken string `json:"scalar_token"`
}

//getScalarToken returns the scalar token for the current user. Used inside the Integration manager.
func getScalarToken(cli *gomatrix.Client, openIDToken *openIDTokenResponse) (resp *scalarRegisterResp, err error) {
	urlPath := "https://scalar.vector.im/api/register"
	_, err = cli.MakeRequest("POST", urlPath, openIDToken, &resp)
	return
}

func ReqAndSaveScalarToken(cli *gomatrix.Client) (err error) {
	cacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Fatalln(DBOpenErr)
	}

	var integToken string

	// Get cache
	DBErr := cacheDB.View(func(txn *badger.Txn) error {
		roomAvatarDataItem, QueryErr := txn.Get([]byte("user|" + cli.UserID + "|integToken"))
		if QueryErr != nil && QueryErr != badger.ErrKeyNotFound {
			return QueryErr
		}
		if QueryErr != badger.ErrKeyNotFound {
			integToken = roomAvatarDataItem.ToString()
			return nil
		}
		return nil
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	if integToken != "" {
		openIdToken, OpenIDerr := getOpenIDToken(cli)
		if OpenIDerr != nil {
			err = OpenIDerr
			return
		}
		var ScalarRegisterErr error
		scalarToken, ScalarRegisterErr := getScalarToken(cli, openIdToken)
		if ScalarRegisterErr != nil {
			err = ScalarRegisterErr
		}
		integToken = scalarToken.ScalarToken

		// Update cache
		DBSetErr := cacheDB.Update(func(txn *badger.Txn) error {
			DBSetErr := txn.Set([]byte("user|"+cli.UserID+"|integToken"), []byte(integToken))
			return DBSetErr
		})
		if DBSetErr != nil {
			err = DBSetErr
			return
		}
	}

	return
}
