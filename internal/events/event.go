package events

import "time"

// EventType represents what happened
type EventType string

const (
	FileMetadataChanged EventType = "FILE_METADATA_CHANGED"
	ConflictDetected    EventType = "CONFLICT_DETECTED"
	ConflictResolved    EventType = "CONFLICT_RESOLVED"
)

// EventSource represents where the change originated
type EventSource string

const (
	SourceDrive EventSource = "GOOGLE_DRIVE"
	SourceApp   EventSource = "APPLICATION"
)

// Event is the single contract between components
type Event struct {
	ID        string
	FileID    string
	Type      EventType
	Source    EventSource
	Version   int64
	Timestamp time.Time

	Payload map[string]any
}
