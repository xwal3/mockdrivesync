package events

// Queue is a simple in-memory event broker
type Queue struct {
	ch chan Event
}

// NewQueue creates a new event queue
func NewQueue(bufferSize int) *Queue {
	return &Queue{
		ch: make(chan Event, bufferSize),
	}
}

// Publish sends an event to the queue
func (q *Queue) Publish(event Event) {
	q.ch <- event
}

// Consume returns a read-only channel for workers
func (q *Queue) Consume() <-chan Event {
	return q.ch
}

// Close the channels
func (q *Queue) Close() {
	close(q.ch)
}
