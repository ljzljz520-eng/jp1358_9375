package calendar

import (
	"errors"
	"sort"
	"time"
)

type Slot struct {
	ID         string
	Start, End time.Time
	Capacity   int
	Booked     int
}

func NewSlot(id string, start time.Time, duration time.Duration, capacity int) Slot {
	return Slot{ID: id, Start: start, End: start.Add(duration), Capacity: capacity}
}
func (s Slot) Available() bool                    { return s.Booked < s.Capacity }
func (s Slot) Overlaps(start, end time.Time) bool { return start.Before(s.End) && end.After(s.Start) }
func (s *Slot) Reserve() error {
	if s.Capacity <= 0 {
		return errors.New("invalid capacity")
	}
	if !s.Available() {
		return errors.New("slot full")
	}
	s.Booked++
	return nil
}
func (s *Slot) Release() {
	if s.Booked > 0 {
		s.Booked--
	}
}
func Sort(slots []Slot) []Slot {
	out := append([]Slot{}, slots...)
	sort.Slice(out, func(i, j int) bool { return out[i].Start.Before(out[j].Start) })
	return out
}
func Find(slots []Slot, id string) (Slot, bool) {
	for _, s := range slots {
		if s.ID == id {
			return s, true
		}
	}
	return Slot{}, false
}
func Within(s Slot, t time.Time) bool { return !t.Before(s.Start) && t.Before(s.End) }
func DaySlots(slots []Slot, day time.Time) []Slot {
	var out []Slot
	y, m, d := day.Date()
	for _, s := range slots {
		sy, sm, sd := s.Start.Date()
		if y == sy && m == sm && d == sd {
			out = append(out, s)
		}
	}
	return Sort(out)
}
