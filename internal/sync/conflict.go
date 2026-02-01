package sync

import (
	"time"

	"github.com/xwal3/mockdrivesync/internal/events"
	"github.com/xwal3/mockdrivesync/internal/state"
)

// Conflict represents a detected metadata conflict
type Conflict struct {
	FileID string
	Field  string

	DriveSnapshot state.MetadataSnapshot
	AppSnapshot   state.MetadataSnapshot

	DetectedAt time.Time
}

// DetectConflict compares two snapshots and returns a conflict if found
func DetectConflict(
	driveSnap state.MetadataSnapshot,
	appSnap state.MetadataSnapshot,
) (*Conflict, bool) {

	// Only compare same file
	if driveSnap.FileID != appSnap.FileID {
		return nil, false
	}

	// Simple example: name conflict
	if driveSnap.Name != appSnap.Name &&
		driveSnap.Version == appSnap.Version {

		return &Conflict{
			FileID:        driveSnap.FileID,
			Field:         "name",
			DriveSnapshot: driveSnap,
			AppSnapshot:   appSnap,
			DetectedAt:    time.Now(),
		}, true
	}

	return nil, false
}

// ConflictToEvent converts a conflict into a sync event
func ConflictToEvent(conflict *Conflict) events.Event {
	return events.Event{
		FileID:    conflict.FileID,
		Type:      events.ConflictDetected,
		Source:    events.SourceApp,
		Version:   time.Now().Unix(),
		Timestamp: time.Now(),
		Payload: map[string]any{
			"field": conflict.Field,
			"drive": conflict.DriveSnapshot,
			"app":   conflict.AppSnapshot,
		},
	}
}
