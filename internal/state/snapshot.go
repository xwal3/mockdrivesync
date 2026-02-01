package state

import "time"

// SnapshotSource indicates where this snapshot originated
type SnapshotSource string

const (
	SnapshotFromDrive SnapshotSource = "GOOGLE_DRIVE"
	SnapshotFromApp   SnapshotSource = "APPLICATION"
)

// MetadataSnapshot represents an immutable view of file metadata
type MetadataSnapshot struct {
	FileID string

	Name string
	Tags []string

	Version   int64
	Source    SnapshotSource
	Timestamp time.Time
}
