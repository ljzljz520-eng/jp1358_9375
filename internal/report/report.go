package report

import (
	"fmt"
	"sort"
	"time"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/storage"
)

type Summary struct {
	Records, Profiles, Events, Audits int
	GeneratedAt                       time.Time
}
type Reporter struct{ Store *storage.Store }

func New(s *storage.Store) *Reporter { return &Reporter{Store: s} }
func (r *Reporter) Summary() (Summary, error) {
	a, e := r.Store.Count("records")
	if e != nil {
		return Summary{}, e
	}
	p, e := r.Store.Count("profiles")
	if e != nil {
		return Summary{}, e
	}
	v, e := r.Store.Count("events")
	if e != nil {
		return Summary{}, e
	}
	h, e := r.Store.Count("audits")
	return Summary{a, p, v, h, time.Now().UTC()}, e
}
func (r Summary) String() string {
	return fmt.Sprintf("records=%d profiles=%d events=%d audits=%d", r.Records, r.Profiles, r.Events, r.Audits)
}
func SortAudits(a []model.Audit) []model.Audit {
	out := append([]model.Audit{}, a...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Filter(a []model.Audit, since time.Time) []model.Audit {
	var out []model.Audit
	for _, v := range a {
		if !v.At.Before(since) {
			out = append(out, v)
		}
	}
	return out
}
