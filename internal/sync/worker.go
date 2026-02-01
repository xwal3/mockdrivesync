package sync

import (
	"log"

	"github.com/xwal3/mockdrivesync/internal/events"
)

type Worker struct {
	ID string
}

// Process handles a single event

func (w *Worker) Process(event events.Event) error {
	log.Printf(
		"[worker=%s] processing event file=%s type=%s source=%s version=%d",
		w.ID,
		event.FileID,
		event.Type,
		event.Source,
		event.Version,
	)

	switch event.Type {

	case events.FileMetadataChanged:
		return w.handleMetadataChange(event)

	case events.ConflictDetected:
		log.Printf("[worker=%s] conflict detected for file=%s", w.ID, event.FileID)
		return nil

	case events.ConflictResolved:
		log.Printf("[worker=%s] conflict resolved for file=%s", w.ID, event.FileID)
		return nil

	default:
		log.Printf("[worker=%s] unknown event type: %s", w.ID, event.Type)
		return nil
	}
}

func (w *Worker) handleMetadataChange(event events.Event) error {
	log.Printf(
		"[worker=%s] applying metadata change for file=%s payload=%v",
		w.ID,
		event.FileID,
		event.Payload,
	)

	// Future:
	// - load latest snapshot
	// - compare versions
	// - detect conflicts
	// - emit conflict or apply update

	return nil
}
