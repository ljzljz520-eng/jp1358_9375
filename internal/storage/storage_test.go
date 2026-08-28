package storage

import (
	"path/filepath"
	"testing"
	"yoga.example/studio/internal/model"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := model.NewRecord("r", "A", "a@b", "n")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r"); e != nil {
		t.Fatal(e)
	}
}
