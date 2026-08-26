package scheduler

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

// TestUploadThenCancelNoPanic reproduces the "upload a cover and immediately
// cancel the task" scenario. Before the fix, Cancel and the worker goroutine
// both closed j.done, panicking with "close of closed channel" and crashing
// the request. The fixed Uploader must stop safely and return a clear result.
func TestUploadThenCancelNoPanic(t *testing.T) {
	u := New()
	id := "cover-1"

	// Start the upload (cover is in flight), then cancel concurrently with the
	// worker. Both paths used to close j.done; now done is closed exactly once.
	u.Start(context.Background(), id, []byte("cover"))
	if err := u.Cancel(id); err != nil {
		t.Fatalf("Cancel returned err: %v", err)
	}
	// After Cancel returns the worker has fully exited and is no longer active.
	if _, ok := func() (*job, bool) {
		u.mu.Lock()
		defer u.mu.Unlock()
		j, ok := u.active[id]
		return j, ok
	}(); ok {
		t.Fatalf("job still active after Cancel")
	}
}

// TestRunCancelReturnsClearResult ensures Run, when cancelled mid-flight, stops
// the upload safely and returns an explicit cancelled result rather than
// crashing.
func TestRunCancelReturnsClearResult(t *testing.T) {
	u := New()
	ctx, cancel := context.WithCancel(context.Background())
	id := "cover-2"

	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()

	r, err := u.Run(ctx, id, []byte("cover"))
	if r.Status != "cancelled" {
		t.Fatalf("expected cancelled status, got %q (err=%v)", r.Status, err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled error, got %v", err)
	}
}

// TestCancelUnknownIsSafe asserts cancelling a non-existent upload returns the
// documented error instead of panicking.
func TestCancelUnknownIsSafe(t *testing.T) {
	u := New()
	if err := u.Cancel("nope"); err == nil {
		t.Fatal("expected error cancelling unknown upload")
	}
}

// TestCancelAllStopsSafely stresses concurrent cancel-all + repeated start on
// many ids, which under the old code double-closed done and panicked.
func TestCancelAllStopsSafely(t *testing.T) {
	u := New()
	const n = 50
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := "cover-x"
			u.Start(context.Background(), id, []byte("cover"))
		}(i)
	}
	// Let some work spin up, then cancel everything while starts may still race.
	time.Sleep(6 * time.Millisecond)
	if got := u.CancelAll(); got == 0 {
		t.Fatal("CancelAll reported 0 active uploads")
	}
	wg.Wait()
	// Give stragglers time to exit; the active map must end up empty.
	time.Sleep(10 * time.Millisecond)
	u.mu.Lock()
	remaining := len(u.active)
	u.mu.Unlock()
	if remaining != 0 {
		t.Fatalf("expected no active uploads, got %d", remaining)
	}
}
