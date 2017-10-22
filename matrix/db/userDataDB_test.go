package db

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	"github.com/shibukawa/configdir"
	"github.com/tidwall/buntdb"
)

// OpenUserDB opens or generates the Database file for settings and Cache
func OpenUserDBTest() (db *buntdb.DB, err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/tests/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/tests/data/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	expDB, DBErr := buntdb.Open(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/userData.db")
	if DBErr != nil {
		err = DBErr
		return
	}

	db = expDB
	return
}

var _ = Describe("UserDataDB", func() {
	var db *buntdb.DB
	/*BeforeEach(func() {
		DBErr := db.Update(func(tx *buntdb.Tx) error {
			DeleteErr := tx.DeleteAll()
			return DeleteErr
		})
		if DBErr != nil {
			Fail(DBErr.Error())
		}
	})*/

	Describe("Test User DB", func() {
		Context("opening", func() {
			It("should open the DB", func() {
				var DBOpenErr error
				db, DBOpenErr = OpenUserDBTest()
				if DBOpenErr != nil {
					Fail(DBOpenErr.Error())
				}
			})
		})
	})
})
