package drive

import "time"

// Change represents a single metadata change from Google Drive
type Change struct {
	FileID       string
	Name         string
	Tags         []string
	ModifiedTime time.Time
	Version      int64
}

// Client defines how we talk to Google Drive
// This is intentionally minimal and only mocks.
type Client interface {
	// FetchChanges returns:
	// - a list of changes since the cursor
	// - the next cursor
	// - an error, if any
	FetchChanges(cursor string) ([]Change, string, error)
}
