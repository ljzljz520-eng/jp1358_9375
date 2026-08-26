package service

import (
	"errors"
	"yoga.example/studio/internal/model"
)

func (s *Service) ArchiveEvent(id string) error {
	e, err := s.Store.GetEvent(id)
	if err != nil {
		return err
	}
	if !e.IsActive() {
		return errors.New("event already archived")
	}
	e = e.Archive()
	if err = s.Store.SaveEvent(e); err != nil {
		return err
	}
	return s.Store.TouchAudit("Event", "archived", id)
}
func (s *Service) ConfirmEvent(id string) error {
	e, err := s.Store.GetEvent(id)
	if err != nil {
		return err
	}
	if e.Status == "archived" {
		return errors.New("archived event")
	}
	e = e.Confirm()
	return s.Store.SaveEvent(e)
}
func EventCopy(e model.Event) model.Event { return e }
