package drive

import "time"

// MockClient simulates Google Drive Changes API
type MockClient struct{}

// FetchChanges returns fake metadata changes and a next cursor
func (m *MockClient) FetchChanges(cursor string) ([]Change, string, error) {

	changes := []Change{
		{
			FileID:       "file-001",
			Name:         "video_v2.mp4",
			Tags:         []string{"review"},
			ModifiedTime: time.Now(),
			Version:      1,
		},
		{
			FileID:       "file-002",
			Name:         "thumbnail_final.png",
			Tags:         []string{"final"},
			ModifiedTime: time.Now(),
			Version:      1,
		},
	}

	nextCursor := "mock-cursor-001"

	return changes, nextCursor, nil
}
