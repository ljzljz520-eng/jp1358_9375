package calendar

import (
	"testing"
	"time"
)

func TestSlotRules(t *testing.T) {
	s := NewSlot("x", time.Now(), time.Hour, 1)
	if !s.Available() {
		t.Fatal()
	}
	if e := s.Reserve(); e != nil || s.Available() {
		t.Fatal(e)
	}
}
