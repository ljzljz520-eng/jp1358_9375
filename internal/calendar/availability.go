package calendar

import (
	"time"
)

func Morning(start time.Time) Slot {
	return NewSlot("morning", start.Truncate(24*time.Hour).Add(8*time.Hour), time.Hour, 8)
}
func Evening(start time.Time) Slot {
	return NewSlot("evening", start.Truncate(24*time.Hour).Add(18*time.Hour), 90*time.Minute, 12)
}
func Normalize(slots []Slot) []Slot {
	out := Sort(slots)
	for i := range out {
		if out[i].Capacity < 1 {
			out[i].Capacity = 1
		}
	}
	return out
}
