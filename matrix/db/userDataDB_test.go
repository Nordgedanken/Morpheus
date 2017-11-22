package db

import (
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger"
	. "github.com/onsi/ginkgo"

	//. "github.com/onsi/gomega"
	"github.com/shibukawa/configdir"
)

// OpenUserDB opens or generates the Database file for settings and Cache
func OpenUserDBTest() (db *badger.DB, err error) {
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
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/user/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/tests/data/user/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/userValue/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/tests/data/userValue/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	opts := badger.DefaultOptions
	opts.Dir = filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/user/"
	opts.ValueDir = filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/tests/data/userValue/"

	expDB, DBErr := badger.Open(opts)
	if DBErr != nil {
		err = DBErr
		return
	}

	db = expDB
	return
}

var _ = Describe("UserDataDB", func() {
	var db *badger.DB
	/*BeforeEach(func() {
		DBErr := db.Update(func(txn *badger.Txn) error {
			DeleteErr := txn.DeleteAll()
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
