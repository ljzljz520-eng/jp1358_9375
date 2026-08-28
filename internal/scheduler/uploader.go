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
type job struct {
	cancel context.CancelFunc
	done   chan struct{}
}
type Uploader struct {
	mu     sync.Mutex
	active map[string]*job
}

func New() *Uploader { return &Uploader{active: make(map[string]*job)} }
func (u *Uploader) Start(ctx context.Context, id string, data []byte) <-chan Result {
	out := make(chan Result, 1)
	work, cancel := context.WithCancel(ctx)
	j := &job{cancel: cancel, done: make(chan struct{})}
	u.mu.Lock()
	u.active[id] = j
	u.mu.Unlock()
	go func() {
		defer func() { close(j.done); u.mu.Lock(); delete(u.active, id); u.mu.Unlock() }()
		select {
		case <-time.After(15 * time.Millisecond):
			out <- Result{ID: id, Status: "uploaded"}
		case <-work.Done():
			out <- Result{ID: id, Status: "cancelled", Err: work.Err()}
		}
	}()
	return out
}
func (u *Uploader) Cancel(id string) error {
	u.mu.Lock()
	j, ok := u.active[id]
	u.mu.Unlock()
	if !ok {
		return errors.New("upload not active")
	}
	j.cancel()
	close(j.done)
	return nil
}
func (u *Uploader) Run(ctx context.Context, id string, data []byte) (Result, error) {
	select {
	case r := <-u.Start(ctx, id, data):
		return r, r.Err
	case <-ctx.Done():
		return Result{ID: id, Status: "cancelled", Err: ctx.Err()}, ctx.Err()
	}
}
func (u *Uploader) CancelAll() int {
	u.mu.Lock()
	defer u.mu.Unlock()
	n := 0
	for _, j := range u.active {
		j.cancel()
		close(j.done)
		n++
	}
	return n
}
func (r Result) Message() string {
	if r.Err != nil {
		return fmt.Sprintf("%s: %v", r.Status, r.Err)
	}
	return r.Status
}
