package main

import (
	"context"
	"path/filepath"
	"testing"
	"time"
	"yoga.example/studio/internal/calendar"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/service"
	"yoga.example/studio/internal/storage"
)

func fixture(t *testing.T) *service.Service {
	t.Helper()
	s, e := storage.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return service.New(s)
}
func TestWorkflowOne(t *testing.T) {
	s := fixture(t)
	if e := s.Register(model.NewRecord("r", "Mia", "mia@example.com", "beginner")); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s := fixture(t)
	s.Register(model.NewRecord("r", "Mia", "mia@example.com", ""))
	s.AddSlot(calendar.NewSlot("s", time.Now().Add(time.Hour), time.Hour, 2))
	if _, e := s.Book("r", "s"); e != nil {
		t.Fatal(e)
	}
	if e := s.CancelBooking("r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s := fixture(t)
	r, e := s.UploadCover(context.Background(), "cover", []byte("bytes"))
	if e != nil || r.Status != "uploaded" {
		t.Fatal(r, e)
	}
}
func TestBusinessChain08(t *testing.T) {
	s := fixture(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan struct{})
	go func() { defer close(done); _, _ = s.UploadCover(ctx, "cover", []byte("cover")) }()
	time.Sleep(time.Millisecond)
	_ = s.CancelUpload("cover")
	<-done
}
