package domain

import "time"

// SessionID uniquely identifies a session across planning, execution, and
// replay operations.
type SessionID string

// Session represents a full xtai interaction including plan, execution, and
// associated artifacts. The structure is intentionally minimal for Phase 1.
type Session struct {
	ID        SessionID
	StartedAt time.Time
	EndedAt   time.Time
	Status    SessionStatus
}

// SessionStatus captures the lifecycle state of a Session.
type SessionStatus string

const (
	SessionStatusPending   SessionStatus = "pending"
	SessionStatusRunning   SessionStatus = "running"
	SessionStatusSucceeded SessionStatus = "succeeded"
	SessionStatusFailed    SessionStatus = "failed"
)

