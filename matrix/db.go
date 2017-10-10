package matrix

import (
	"log"
	"os"
	"path/filepath"

	"github.com/matrix-org/gomatrix"
	"github.com/shibukawa/configdir"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func OpenDB() (expDB *buntdb.DB) {
	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheus")
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

func InitData(cli *gomatrix.Client, db *buntdb.DB) {
	roomsURL := cli.BuildURL("joined_rooms")
	var rooms JoinedRooms
	_, err := cli.MakeRequest("GET", roomsURL, nil, &rooms)
	if err != nil {
		localLog.Println(err)
	}

	for _, room := range rooms.JoinedRooms {
		var roomAliases RoomAliases
		if err := cli.StateEvent(room, "m.room.aliases", "", &roomAliases); err != nil {
			localLog.Println(err)
		}

		for index, alias := range roomAliases.Content.Aliases {
			// Update cache
			DBerr := db.Update(func(tx *buntdb.Tx) error {
				tx.Set("room:"+room+":aliases"+string(index), alias, nil)
				return nil
			})
			if DBerr != nil {
				localLog.Fatalln(DBerr)
			}
		}

	}
}
