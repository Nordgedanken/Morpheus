package db

import (
	"github.com/dgraph-io/badger"
)

// Storer is the interface which needs to be conformed to in order to persist Go-NEB data
type Storer interface {
	UpdateNextBatch(userID, nextBatch string) (err error)
	LoadNextBatch(userID string) (nextBatch string, err error)
}

// MorpheusStorage is the StorageInterface which needs to be conformed to in order to persist Go-NEB data
type MorpheusStorage struct {
	Database *badger.DB
}

// UpdateNextBatch updates the next_batch token for the given user.
func (m *MorpheusStorage) UpdateNextBatch(userID, nextBatch string) (err error) {
	DBerr := m.Database.Update(func(txn *badger.Txn) error {

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
	DBerr := m.Database.View(func(txn *badger.Txn) error {

		nextBatchItem, NextBatchErr := txn.Get([]byte("matrix|" + userID + "|nextBatch|"))
		if NextBatchErr != nil {
			return NextBatchErr
		}

		nextBatchByte, nextBatchByteErr := nextBatchItem.Value()
		if nextBatchByteErr != nil {
			return nextBatchByteErr
		}

		nextBatch = string(nextBatchByte)
		return nil
	})
	if DBerr != nil {
		err = DBerr
		return
	}
	return
}
