package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"strings"
	"yoga.example/studio/internal/model"
)

func (s *Store) SearchRecords(term string) ([]model.Record, error) {
	var out []model.Record
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["records"]).ForEach(func(_, v []byte) error {
			var r model.Record
			if err := json.Unmarshal(v, &r); err != nil {
				return err
			}
			if term == "" || strings.Contains(strings.ToLower(r.Name), strings.ToLower(term)) || strings.Contains(strings.ToLower(r.Email), strings.ToLower(term)) {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets["records"]).Delete([]byte(id)) })
}
func (s *Store) SaveMany(records []model.Record) error {
	for _, r := range records {
		if err := s.SaveRecord(r); err != nil {
			return err
		}
	}
	return nil
}
