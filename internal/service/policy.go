package service

import (
	"errors"
	"strings"
	"time"
	"yoga.example/studio/internal/calendar"
	"yoga.example/studio/internal/model"
)

func NormalizeName(v string) string { return strings.Join(strings.Fields(strings.TrimSpace(v)), " ") }
func CanBook(r model.Record, slot calendar.Slot, now time.Time) error {
	if e := r.Validate(); e != nil {
		return e
	}
	if !slot.Available() {
		return errors.New("capacity exhausted")
	}
	if slot.Start.Before(now) {
		return errors.New("slot in past")
	}
	return nil
}
func StatusMessage(status string) string {
	switch status {
	case "uploaded":
		return "Cover is ready"
	case "cancelled":
		return "Upload cancelled"
	default:
		return "Upload pending"
	}
}
func IsWeekend(t time.Time) bool { return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday }
func ClampCapacity(v int) int {
	if v < 1 {
		return 1
	}
	if v > 100 {
		return 100
	}
	return v
}
