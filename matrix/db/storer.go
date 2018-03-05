package db

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

// Storer is the interface which needs to be conformed to in order to persist Go-NEB data
type Storer interface {
	UpdateNextBatch(userID, nextBatch string) (err error)
	LoadNextBatch(userID string) (nextBatch string, err error)
}

// MorpheusStorage is the StorageInterface which needs to be conformed to in order to persist Go-NEB data
type MorpheusStorage struct {
}

// UpdateNextBatch updates the next_batch token for the given user.
func (m *MorpheusStorage) UpdateNextBatch(userID, nextBatch string) (err error) {
	CacheDB, DBOpenErr := OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}
	DBerr := CacheDB.Update(func(txn *badger.Txn) error {

		DBSetNextBatchErr := txn.Set([]byte("matrix|"+userID+"|nextBatch|"), []byte(nextBatch))
		return DBSetNextBatchErr
	})
	if DBerr != nil {
		err = DBerr
		return
	}
	return
}

// LoadNextBatch loads the next_batch token for the given user.
func (m *MorpheusStorage) LoadNextBatch(userID string) (nextBatch string, err error) {
	CacheDB, DBOpenErr := OpenCacheDB()
	if DBOpenErr != nil {
		err = DBOpenErr
		return
	}
	DBerr := CacheDB.View(func(txn *badger.Txn) error {
		nextBatchResult, QueryErr := Get(txn, []byte("user|accessToken"))
		if QueryErr != nil {
			return QueryErr
		}
		nextBatch = fmt.Sprintf("%s", nextBatchResult)
		return nil
	})
	if DBerr != nil {
		err = DBerr
		return
	}
	return
}
