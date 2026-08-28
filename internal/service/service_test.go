package service

import (
	"path/filepath"
	"testing"
	"time"
	"yoga.example/studio/internal/calendar"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/storage"
)

func TestRegisterBook(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer st.Close()
	s := New(st)
	if e := s.Register(model.NewRecord("r", "A", "a@b", "")); e != nil {
		t.Fatal(e)
	}
	slot := calendar.NewSlot("s", time.Now().Add(time.Hour), time.Hour, 1)
	s.AddSlot(slot)
	if _, e := s.Book("r", "s"); e != nil {
		t.Fatal(e)
	}
}
