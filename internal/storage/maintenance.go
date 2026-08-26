package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"sort"
	"time"
	"yoga.example/studio/internal/model"
)

func (s *Store) ListRecords() ([]model.Record, error) {
	var out []model.Record
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["records"]).ForEach(func(_, v []byte) error {
			var r model.Record
			if err := decode(v, &r); err != nil {
				return err
			}
			out = append(out, r)
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, e
}
func decode(data []byte, v any) error        { return jsonUnmarshal(data, v) }
func jsonUnmarshal(data []byte, v any) error { return unmarshal(data, v) }
func unmarshal(data []byte, v any) error     { return json.Unmarshal(data, v) }
func (s *Store) PurgeAudits(before time.Time) (int, error) {
	removed := 0
	e := s.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["audits"]).ForEach(func(k, v []byte) error {
			var a model.Audit
			if err := decode(v, &a); err != nil {
				return err
			}
			if a.At.Before(before) {
				if err := tx.Bucket(buckets["audits"]).Delete(k); err != nil {
					return err
				}
				removed++
			}
			return nil
		})
	})
	return removed, e
}
