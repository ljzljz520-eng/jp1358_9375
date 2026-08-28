package scheduler

import (
	"context"
	"testing"
	"time"
)

func TestUploadCompletes(t *testing.T) {
	u := New()
	r := <-u.Start(context.Background(), "x", []byte("a"))
	if r.Status != "uploaded" {
		t.Fatal(r)
	}
}
func TestCancel(t *testing.T) {
	u := New()
	u.Start(context.Background(), "x", []byte("a"))
	time.Sleep(time.Millisecond)
	if e := u.Cancel("x"); e != nil {
		t.Fatal(e)
	}
}
