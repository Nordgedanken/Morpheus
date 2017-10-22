package db

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/shibukawa/configdir"
	"github.com/tidwall/buntdb"
)

// OpenCacheDB opens or generates the Database file for settings and Cache
func OpenCacheDB() (db *buntdb.DB, err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/data/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	expDB, DBErr := buntdb.Open(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/cache.db")
	if DBErr != nil {
		err = DBErr
		return
	}

	db = expDB
	return
}

// OpenUserDB opens or generates the Database file for settings and Cache
func OpenUserDB() (db *buntdb.DB, err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/data/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	expDB, DBErr := buntdb.Open(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/userData.db")
	if DBErr != nil {
		err = DBErr
		return
	}

	db = expDB
	return
}

// CacheMessageEvents writes message infos into the cache into the defined room
func CacheMessageEvents(id, sender, roomID, message string, timestamp int64) (err error) {
	db, DBOpenErr := OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}
	defer db.Close()

	// Update cache
	DBerr := db.Update(func(tx *buntdb.Tx) error {
		_, _, DBSetIDErr := tx.Set("room|"+roomID+"|messages|"+id+"|id", id, nil)
		if DBSetIDErr != nil {
			return DBSetIDErr
		}

		_, _, DBSetSenderErr := tx.Set("room|"+roomID+"|messages|"+id+"|sender", sender, nil)
		if DBSetSenderErr != nil {
			return DBSetSenderErr
		}

		_, _, DBSetMessageErr := tx.Set("room|"+roomID+"|messages|"+id+"|message", message, nil)
		if DBSetMessageErr != nil {
			return DBSetMessageErr
		}

		timestampString := strconv.FormatInt(timestamp, 10)
		_, _, DBSeTimestampErr := tx.Set("room|"+roomID+"|messages|"+id+"|timestamp", timestampString, nil)
		return DBSeTimestampErr

	})
	if DBerr != nil {
		err = DBerr
		return
	}
	return
}
