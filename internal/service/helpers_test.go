package service

import (
	"path/filepath"
	"testing"

	"yoga.example/studio/internal/storage"
)

// newTestStore opens an isolated bbolt store in the test's temp directory so
// upload/cancel tests never touch a real database file.
func newTestStore(t *testing.T) *storage.Store {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	st, err := storage.Open(path)
	if err != nil {
		t.Fatalf("open test store: %v", err)
	}
	return st
}
