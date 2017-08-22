package matrix

import (
	"log"
	"os"

	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func OpenDB() (expDB *buntdb.DB) {
	// Open the data.db file. It will be created if it doesn't exist.
	if _, err := os.Stat("./data/"); os.IsNotExist(err) {
		os.Mkdir("./data/", 0666)
	}
	expDB, err := buntdb.Open("./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	db = expDB

	return
}
