package model

import (
	"testing"
	"time"
)

func TestRecordValidation(t *testing.T) {
	if NewRecord("1", "A", "a@b", " ").Validate() != nil {
		t.Fatal("valid record rejected")
	}
	if (Record{}).Validate() == nil {
		t.Fatal("invalid accepted")
	}
}
func TestEventLifecycle(t *testing.T) {
	e := NewEvent("e", "r", "class", time.Now())
	if !e.Confirm().IsActive() {
		t.Fatal("not active")
	}
	if e.Archive().IsActive() {
		t.Fatal("archive active")
	}
}
