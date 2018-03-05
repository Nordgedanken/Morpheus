package syncer

import (
	"fmt"

	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/dgraph-io/badger"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
)

type Storer interface {
	SaveFilterID(userID, filterID string)
	LoadFilterID(userID string) string
	SaveNextBatch(userID, nextBatchToken string)
	LoadNextBatch(userID string) string
	SaveRoom(room *gomatrix.Room)
	LoadRoom(roomID string) *gomatrix.Room
}

// MorpheusStore implements the gomatrix.Storer interface.
//
// It persists the next batch token in the database, and includes a ClientConfig for the client.
type MorpheusStore struct {
	gomatrix.Storer
	cli *gomatrix.Client
}

// NewInMemoryStore constructs a new MorpheusStorage.
func NewMorpheusStore(cli *gomatrix.Client) *MorpheusStore {
	return &MorpheusStore{
		cli: cli,
	}
}

// SaveFilterID to memory.
func (m *MorpheusStore) SaveFilterID(userID, filterID string) {
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return
	}
	DBerr := CacheDB.Update(func(txn *badger.Txn) error {
		DBSetFilterIDErr := txn.Set([]byte("matrix|"+userID+"|filterID|"), []byte(filterID))
		return DBSetFilterIDErr
	})
	if DBerr != nil {
		log.Errorln(DBerr)
		return
	}
	return
}

// LoadFilterID from memory.
func (m *MorpheusStore) LoadFilterID(userID string) string {
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return ""
	}
	var filterID string
	DBerr := CacheDB.View(func(txn *badger.Txn) error {
		filterIDResult, QueryErr := db.Get(txn, []byte("user|filterID"))
		if QueryErr != nil {
			return QueryErr
		}
		filterID = fmt.Sprintf("%s", filterIDResult)
		return nil
	})
	if DBerr != nil {
		log.Errorln(DBerr)
		return ""
	}
	return filterID
}

// SaveNextBatch saves to the database.
func (m *MorpheusStore) SaveNextBatch(userID, nextBatch string) {
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return
	}
	DBerr := CacheDB.Update(func(txn *badger.Txn) error {
		DBSetNextBatchErr := txn.Set([]byte("matrix|"+userID+"|nextBatch|"), []byte(nextBatch))
		return DBSetNextBatchErr
	})
	if DBerr != nil {
		log.Errorln(DBerr)
		return
	}
	return
}

// LoadNextBatch loads from the database.
func (m *MorpheusStore) LoadNextBatch(userID string) string {
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return ""
	}
	var nextBatch string
	DBerr := CacheDB.View(func(txn *badger.Txn) error {
		nextBatchResult, QueryErr := db.Get(txn, []byte("user|accessToken"))
		if QueryErr != nil {
			return QueryErr
		}
		nextBatch = fmt.Sprintf("%s", nextBatchResult)
		return nil
	})
	if DBerr != nil {
		log.Errorln(DBerr)
		return ""
	}
	return nextBatch
}

// SaveRoom to memory.
func (m *MorpheusStore) SaveRoom(room *gomatrix.Room) {
	CacheDB, DBOpenErr := db.OpenCacheDB()
	if DBOpenErr != nil {
		log.Errorln(DBOpenErr)
		return
	}
	roomID := room.ID

	DBerr := CacheDB.Update(func(txn *badger.Txn) error {

		DBSetRoomNameErr := txn.Set([]byte("room|"+roomID+"|id"), []byte(roomID))
		return DBSetRoomNameErr
	})
	if DBerr != nil {
		log.Errorln(DBerr)
		return
	}
	return
}

// LoadRoom from memory.
func (m *MorpheusStore) LoadRoom(roomID string) *gomatrix.Room {
	room := gomatrix.NewRoom(roomID)
	return room
}
