package storage

import "context"

// SessionStore is a placeholder for session persistence. A concrete
// implementation using SQLite or JSONL will be added separately.
type SessionStore struct{}

// Save is a stub method that will persist sessions once storage is implemented.
func (s *SessionStore) Save(ctx context.Context) error {
	_ = ctx
	return nil
}

