package main

import (
	"path/filepath"
	"testing"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/storage"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "persist.db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveRecord(model.NewRecord("persist", "P", "p@example.com", "")); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
