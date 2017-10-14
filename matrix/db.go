package matrix

import (
	"os"
	"path/filepath"

	"github.com/matrix-org/gomatrix"
	"github.com/shibukawa/configdir"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

// OpenDB opens or generates the Database file for settings and Cache
func OpenDB() (db *buntdb.DB, err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheus")
	if _, StatErr := os.Stat(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/"); os.IsNotExist(StatErr) {
		MkdirErr := os.MkdirAll(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)+"/data/", 0666)
		if MkdirErr != nil {
			err = MkdirErr
			return
		}
	}
	expDB, DBErr := buntdb.Open(filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path) + "/data/data.db")
	if DBErr != nil {
		err = DBErr
		return
	}

	db = expDB
	return
}

// InitData inits basic Data like getting aliases of joinedRooms
func InitData(cli *gomatrix.Client) (err error) {
	db, DBOpenErr := OpenDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}
	defer db.Close()

	roomsURL := cli.BuildURL("joined_rooms")
	var rooms JoinedRooms
	_, ReqErr := cli.MakeRequest("GET", roomsURL, nil, &rooms)
	if ReqErr != nil {
		err = ReqErr
		return
	}

	for _, room := range rooms.JoinedRooms {
		var roomAliases RoomAliases
		if StateEventErr := cli.StateEvent(room, "m.room.aliases", "", &roomAliases); StateEventErr != nil {
			localLog.Println(StateEventErr)
			// Not returning as a Error NotFound is allowed
		}

		for index, alias := range roomAliases.Content.Aliases {
			// Update cache
			DBerr := db.Update(func(tx *buntdb.Tx) error {
				localLog.Println(room)
				_, _, DBSetErr := tx.Set("room:"+room+":aliases:"+string(index), alias, nil)
				return DBSetErr
			})
			if DBerr != nil {
				err = DBerr
				return
			}
		}

	}
	return
}

// CacheMessageEvents writes message infos into the cache into the defined room
func CacheMessageEvents(id, sender, roomID, message string, timestamp int64) (err error) {
	db, DBOpenErr := OpenDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}
	defer db.Close()

	// Update cache
	DBerr := db.Update(func(tx *buntdb.Tx) error {
		localLog.Println(roomID)
		_, _, DBSetIDErr := tx.Set("room:"+roomID+":messages:"+id+":id", id, nil)
		if DBSetIDErr != nil {
			return DBSetIDErr
		}

		_, _, DBSetSenderErr := tx.Set("room:"+roomID+":messages:"+id+":sender", sender, nil)
		if DBSetSenderErr != nil {
			return DBSetSenderErr
		}

		_, _, DBSetMessageErr := tx.Set("room:"+roomID+":messages:"+id+":message", message, nil)
		if DBSetMessageErr != nil {
			return DBSetMessageErr
		}

		_, _, DBSeTimestampErr := tx.Set("room:"+roomID+":messages:"+id+":timestamp", string(timestamp), nil)
		return DBSeTimestampErr

	})
	if DBerr != nil {
		err = DBerr
		return
	}
	return
}
