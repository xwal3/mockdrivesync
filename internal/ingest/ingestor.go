package ingest

import (
	"time"

	"github.com/xwal3/mockdrivesync/internal/drive"
	"github.com/xwal3/mockdrivesync/internal/events"
)

type Ingestor struct {
	DriveClient drive.Client
	EventQueue  *events.Queue
}

// Run performs one incremental sync cycle
func (i *Ingestor) Run(cursor string) (string, error) {

	changes, nextCursor, err := i.DriveClient.FetchChanges(cursor)

	if err != nil {
		return cursor, err
	}

	for _, change := range changes {
		event := events.Event{
			FileID:    change.FileID,
			Type:      events.FileMetadataChanged,
			Source:    events.SourceDrive,
			Version:   change.Version,
			Timestamp: time.Now(),
			Payload: map[string]any{
				"name": change.Name,
				"tags": change.Tags,
			},
		}

		i.EventQueue.Publish(event)
	}

	return nextCursor, nil
}
