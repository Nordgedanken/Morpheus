package syncer

import (
	"github.com/Nordgedanken/Morpheus/matrix/db"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
)

// MorpheusStore implements the gomatrix.Storer interface.
//
// It persists the next batch token in the database, and includes a ClientConfig for the client.
type MorpheusStore struct {
	gomatrix.InMemoryStore
	CacheDatabase db.Storer
}

// SaveNextBatch saves to the database.
func (m *MorpheusStore) SaveNextBatch(userID, nextBatch string) {
	if err := m.CacheDatabase.UpdateNextBatch(userID, nextBatch); err != nil {
		log.WithFields(log.Fields{
			log.ErrorKey: err,
			"user_id":    userID,
			"next_batch": nextBatch,
		}).Error("Failed to persist next_batch token")
	}
}

// LoadNextBatch loads from the database.
func (m *MorpheusStore) LoadNextBatch(userID string) string {
	token, err := m.CacheDatabase.LoadNextBatch(userID)
	if err != nil {
		log.WithFields(log.Fields{
			log.ErrorKey: err,
			"user_id":    userID,
		}).Error("Failed to load next_batch token")
		return ""
	}
	return token
}
