package matrix

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shibukawa/configdir"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func OpenDB() (expDB *buntdb.DB) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Neo")
	if _, err := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/"); os.IsNotExist(err) {
		os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/data/", 0666)
	}
	expDB, err := buntdb.Open(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	db = expDB

	return
}
