package scheduler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	ID     string
	Status string
	Err    error
}

// job tracks a single in-flight upload. The done channel is closed exactly
// once (guarded by once) when the worker goroutine exits; that close is the
// single signal that the job has safely stopped.
type job struct {
	id     string
	cancel context.CancelFunc
	done   chan struct{}
	once   sync.Once
}

type Uploader struct {
	mu     sync.Mutex
	active map[string]*job
}

func New() *Uploader { return &Uploader{active: make(map[string]*job)} }

// Start begins an upload for id and returns a buffered channel that receives
// exactly one Result: "uploaded" once the work completes, or "cancelled" if
// the work context is cancelled first.
func (u *Uploader) Start(ctx context.Context, id string, data []byte) <-chan Result {
	out := make(chan Result, 1)
	work, cancel := context.WithCancel(ctx)
	j := &job{id: id, cancel: cancel, done: make(chan struct{})}
	u.mu.Lock()
	u.active[id] = j
	u.mu.Unlock()
	go func() {
		// Release the work context resources, then signal a safe stop. done is
		// closed exactly once so concurrent Cancel calls never double-close it.
		defer func() {
			cancel()
			j.once.Do(func() { close(j.done) })
			u.remove(j)
		}()
		select {
		case <-time.After(15 * time.Millisecond):
			out <- Result{ID: id, Status: "uploaded"}
		case <-work.Done():
			out <- Result{ID: id, Status: "cancelled", Err: work.Err()}
		}
	}()
	return out
}

// Cancel requests a safe stop of an in-flight upload: it signals the worker to
// stop, waits until it has fully exited, and only then returns. It is safe to
// call concurrently with Start, Run, CancelAll, and other Cancel calls; done
// is closed exactly once so there is no double-close panic.
func (u *Uploader) Cancel(id string) error {
	u.mu.Lock()
	j, ok := u.active[id]
	u.mu.Unlock()
	if !ok {
		return errors.New("upload not active")
	}
	j.cancel()
	<-j.done
	u.remove(j)
	return nil
}

// Run starts an upload and blocks until it produces a result or ctx is
// cancelled. On cancellation the in-flight upload is stopped safely before Run
// returns a clear cancelled result.
func (u *Uploader) Run(ctx context.Context, id string, data []byte) (Result, error) {
	out := u.Start(ctx, id, data)
	select {
	case r := <-out:
		return r, r.Err
	case <-ctx.Done():
		// Stop the worker safely. It emits exactly one result before closing
		// done, so after Cancel returns that result is waiting on out.
		_ = u.Cancel(id)
		r := <-out
		return r, r.Err
	}
}

// CancelAll stops every in-flight upload safely and returns the count.
func (u *Uploader) CancelAll() int {
	u.mu.Lock()
	jobs := make([]*job, 0, len(u.active))
	for _, j := range u.active {
		jobs = append(jobs, j)
	}
	u.mu.Unlock()
	for _, j := range jobs {
		j.cancel()
		<-j.done
		u.remove(j)
	}
	return len(jobs)
}

// remove deletes the job from the active set iff it still owns its id, so a
// stale goroutine can never evict a newer job that reused the same id.
func (u *Uploader) remove(j *job) {
	u.mu.Lock()
	if u.active[j.id] == j {
		delete(u.active, j.id)
	}
	u.mu.Unlock()
}

func (r Result) Message() string {
	if r.Err != nil {
		return fmt.Sprintf("%s: %v", r.Status, r.Err)
	}
	return r.Status
}
