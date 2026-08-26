package model

import (
	"fmt"
	"time"
)

func (r Record) Label() string { return fmt.Sprintf("%s <%s>", r.Name, r.Email) }
func (p Profile) Summary() string {
	return fmt.Sprintf("%s: %d specialties", p.Bio, len(p.Specialties))
}
func (e Event) DateKey() string             { return e.Start.UTC().Format("2006-01-02") }
func (a Audit) IsRecent(now time.Time) bool { return now.Sub(a.At) < 24*time.Hour }
