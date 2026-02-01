package ingest

import (
	"testing"

	"github.com/xwal3/mockdrivesync/internal/drive"
	"github.com/xwal3/mockdrivesync/internal/events"
)

func TestIngestorEmitsEvents(t *testing.T) {
	queue := events.NewQueue(10)
	client := &drive.MockClient{}

	ingestor := Ingestor{
		DriveClient: client,
		EventQueue:  queue,
	}

	_, err := ingestor.Run("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	select {
	case <-queue.Consume():
		// success
	default:
		t.Fatal("expected event to be published")
	}
}
