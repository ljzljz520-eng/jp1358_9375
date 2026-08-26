package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"yoga.example/studio/internal/calendar"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/scheduler"
	"yoga.example/studio/internal/storage"
)

type Service struct {
	Store    *storage.Store
	Uploads  *scheduler.Uploader
	mu       sync.Mutex
	slots    map[string]calendar.Slot
	bookings map[string]string
}

func New(st *storage.Store) *Service {
	return &Service{Store: st, Uploads: scheduler.New(), slots: make(map[string]calendar.Slot), bookings: make(map[string]string)}
}
func (s *Service) Register(r model.Record) error {
	if e := r.Validate(); e != nil {
		return e
	}
	if e := s.Store.SaveRecord(r); e != nil {
		return e
	}
	return s.Store.TouchAudit("Record", "registered", r.ID)
}
func (s *Service) Profile(p model.Profile) error {
	if e := p.Validate(); e != nil {
		return e
	}
	if e := s.Store.SaveProfile(p); e != nil {
		return e
	}
	return s.Store.TouchAudit("Profile", "published", p.ID)
}
func (s *Service) AddSlot(slot calendar.Slot) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if slot.ID == "" || slot.End.Before(slot.Start) {
		return errors.New("invalid slot")
	}
	s.slots[slot.ID] = slot
	return nil
}
func (s *Service) Book(recordID, slotID string) (model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, e := s.Store.GetRecord(recordID)
	if e != nil {
		return model.Event{}, e
	}
	if e = r.Validate(); e != nil {
		return model.Event{}, e
	}
	slot, ok := s.slots[slotID]
	if !ok {
		return model.Event{}, errors.New("slot missing")
	}
	if !slot.Available() {
		return model.Event{}, errors.New("slot full")
	}
	if _, dup := s.bookings[recordID]; dup {
		return model.Event{}, errors.New("already booked")
	}
	if e = slot.Reserve(); e != nil {
		return model.Event{}, e
	}
	s.slots[slotID] = slot
	ev := model.NewEvent(fmt.Sprintf("booking-%d", time.Now().UnixNano()), recordID, "class", slot.Start).Confirm()
	s.bookings[recordID] = ev.ID
	if err := s.Store.SaveEvent(ev); err != nil {
		return model.Event{}, err
	}
	return ev, s.Store.TouchAudit("Event", "booked", ev.ID)
}
func (s *Service) CancelBooking(recordID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.bookings[recordID]; !ok {
		return errors.New("booking missing")
	}
	delete(s.bookings, recordID)
	return s.Store.TouchAudit("Event", "cancelled", recordID)
}
func (s *Service) UploadCover(ctx context.Context, id string, data []byte) (scheduler.Result, error) {
	if len(data) == 0 {
		return scheduler.Result{}, errors.New("empty cover")
	}
	result, err := s.Uploads.Run(ctx, id, data)
	_ = s.Store.TouchAudit("Event", "upload_"+result.Status, id)
	return result, err
}
func (s *Service) CancelUpload(id string) error                 { return s.Uploads.Cancel(id) }
func (s *Service) History(entity string) ([]model.Audit, error) { return s.Store.ListAudits(entity) }
