package calendar

import (
	"errors"
	"time"
)

func ValidateWindow(start, end time.Time) error {
	if start.IsZero() || end.IsZero() {
		return errors.New("time required")
	}
	if !end.After(start) {
		return errors.New("end after start")
	}
	if end.Sub(start) > 4*time.Hour {
		return errors.New("window too long")
	}
	return nil
}
func BusinessHours(start, end time.Time) bool {
	if start.Hour() < 6 || end.Hour() > 22 {
		return false
	}
	return start.Weekday() != time.Sunday
}
func NextAvailable(slots []Slot, after time.Time) (Slot, bool) {
	for _, s := range Sort(slots) {
		if s.Start.After(after) && s.Available() {
			return s, true
		}
	}
	return Slot{}, false
}
