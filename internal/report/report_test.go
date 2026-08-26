package report

import (
	"path/filepath"
	"testing"
	"yoga.example/studio/internal/storage"
)

func TestSummary(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	x, e := New(s).Summary()
	if e != nil || x.Records != 0 {
		t.Fatal(x, e)
	}
}
