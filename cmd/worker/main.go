package main

import (
	"log"
	s "sync"

	"github.com/xwal3/mockdrivesync/internal/drive"
	"github.com/xwal3/mockdrivesync/internal/events"
	"github.com/xwal3/mockdrivesync/internal/ingest"
	"github.com/xwal3/mockdrivesync/internal/sync"
)

func main() {
	log.Println("starting mockdrivesync worker")

	queue := events.NewQueue(10)
	driveClient := &drive.MockClient{}

	ingestor := ingest.Ingestor{
		DriveClient: driveClient,
		EventQueue:  queue,
	}

	worker := &sync.Worker{ID: "worker-1"}

	cursor := ""
	nextCursor, err := ingestor.Run(cursor)
	if err != nil {
		log.Fatalf("ingestion failed: %v", err)
	}
	log.Printf("ingestion complete, next cursor=%s", nextCursor)

	queue.Close()

	var wg s.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for event := range queue.Consume() {
			if err := worker.Process(event); err != nil {
				log.Printf("failed to process event: %v", err)
			}
		}
	}()

	wg.Wait()
	log.Println("mockdrivesync finished cleanly")
}
