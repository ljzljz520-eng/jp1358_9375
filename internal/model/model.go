package model

import (
	"errors"
	"strings"
	"time"
)

type Record struct {
	ID, Name, Email, Notes string
	CreatedAt              time.Time
}
type Profile struct {
	ID, Bio                   string
	Specialties, Availability []string
}
type Event struct {
	ID, RecordID, Kind string
	Start              time.Time
	Status             string
}
type Audit struct {
	ID, Entity, Action string
	At                 time.Time
	Detail             string
}

func (r Record) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name required")
	}
	if !strings.Contains(r.Email, "@") {
		return errors.New("valid email required")
	}
	return nil
}
func (p Profile) Validate() error {
	if p.ID == "" || strings.TrimSpace(p.Bio) == "" {
		return errors.New("profile identity and bio required")
	}
	if len(p.Specialties) == 0 {
		return errors.New("specialty required")
	}
	return nil
}
func (e Event) Validate() error {
	if e.ID == "" || e.Kind == "" || e.Start.IsZero() {
		return errors.New("event fields required")
	}
	if e.Status == "" {
		return errors.New("event status required")
	}
	return nil
}
func (a Audit) Validate() error {
	if a.ID == "" || a.Entity == "" || a.Action == "" || a.At.IsZero() {
		return errors.New("audit fields required")
	}
	return nil
}
func NewRecord(id, name, email, notes string) Record {
	return Record{ID: id, Name: name, Email: email, Notes: notes, CreatedAt: time.Now().UTC()}
}
func NewProfile(id, bio string, specialties, availability []string) Profile {
	return Profile{ID: id, Bio: bio, Specialties: append([]string{}, specialties...), Availability: append([]string{}, availability...)}
}
func NewEvent(id, record, kind string, start time.Time) Event {
	return Event{ID: id, RecordID: record, Kind: kind, Start: start, Status: "pending"}
}
func NewAudit(id, entity, action, detail string) Audit {
	return Audit{ID: id, Entity: entity, Action: action, At: time.Now().UTC(), Detail: detail}
}
func (e Event) IsActive() bool { return e.Status == "pending" || e.Status == "confirmed" }
func (e Event) Confirm() Event { e.Status = "confirmed"; return e }
func (e Event) Archive() Event { e.Status = "archived"; return e }
