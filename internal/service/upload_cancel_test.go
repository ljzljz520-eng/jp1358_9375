package service

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

// TestUploadCoverThenCancel reproduces the reported scenario: a cover upload
// is started and the task is cancelled while it is in flight. The service must
// stop the upload safely and record a clear cancelled result instead of
// crashing the request.
func TestUploadCoverThenCancel(t *testing.T) {
	st := newTestStore(t)
	defer st.Close()
	svc := New(st)
	id := "rec-cover-1"

	// Start the cover upload (it stays in flight for ~15ms), then cancel it
	// mid-flight from another goroutine while UploadCover is still running.
	var (
		resMu sync.Mutex
		res   struct {
			status string
			err    error
		}
		done = make(chan struct{})
	)
	go func() {
		defer close(done)
		r, e := svc.UploadCover(context.Background(), id, []byte("cover"))
		resMu.Lock()
		res.status = r.Status
		res.err = e
		resMu.Unlock()
	}()
	// Give the worker a moment to register, then cancel mid-flight.
	time.Sleep(2 * time.Millisecond)
	if err := svc.CancelUpload(id); err != nil {
		t.Fatalf("CancelUpload returned err: %v", err)
	}
	// Wait for the in-flight UploadCover to finish writing its audit.
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("UploadCover did not return after Cancel")
	}

	resMu.Lock()
	status, err := res.status, res.err
	resMu.Unlock()
	if status != "cancelled" {
		t.Fatalf("expected cancelled result, got %q (err=%v)", status, err)
	}

	// The audit trail must reflect the cancelled upload, not a crash.
	audits, err := svc.History("Event")
	if err != nil {
		t.Fatalf("History err: %v", err)
	}
	var sawCancelled bool
	for _, a := range audits {
		if a.Action == "upload_cancelled" {
			sawCancelled = true
		}
	}
	if !sawCancelled {
		t.Fatalf("expected upload_cancelled audit, got %+v", audits)
	}
}

// TestUploadCoverCancelUnknown confirms cancelling an unknown upload surfaces a
// clear error rather than panicking.
func TestUploadCoverCancelUnknown(t *testing.T) {
	st := newTestStore(t)
	defer st.Close()
	svc := New(st)
	if err := svc.CancelUpload("missing"); err == nil {
		t.Fatal("expected error cancelling unknown upload")
	}
}

// TestUploadCoverCancelledContext ensures that when the caller's context is
// cancelled mid-upload, UploadCover returns a clear cancelled result.
func TestUploadCoverCancelledContext(t *testing.T) {
	st := newTestStore(t)
	defer st.Close()
	svc := New(st)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Millisecond)
		cancel()
	}()
	r, err := svc.UploadCover(ctx, "rec-cover-ctx", []byte("cover"))
	if r.Status != "cancelled" {
		t.Fatalf("expected cancelled status, got %q (err=%v)", r.Status, err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}
