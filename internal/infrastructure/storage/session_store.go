package storage

import "context"

// SessionStore is a placeholder for session persistence. A concrete
// implementation using SQLite or JSONL will be added in a later phase.
type SessionStore struct{}

// Save is a stub method that will persist sessions in a future phase.
func (s *SessionStore) Save(ctx context.Context) error {
	_ = ctx
	return nil
}

