package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"path/filepath"
	"time"
	"yoga.example/studio/internal/model"
)

var buckets = map[string][]byte{"records": []byte("records"), "profiles": []byte("profiles"), "events": []byte("events"), "audits": []byte("audits")}

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, bbolt.DefaultOptions)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) Path() string { return s.path }
func put[T any](s *Store, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](s *Store, b []byte, key string) (T, error) {
	var out T
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return errors.New("not found")
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func (s *Store) SaveRecord(r model.Record) error { return put(s, buckets["records"], r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	return get[model.Record](s, buckets["records"], id)
}
func (s *Store) SaveProfile(p model.Profile) error { return put(s, buckets["profiles"], p.ID, p) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	return get[model.Profile](s, buckets["profiles"], id)
}
func (s *Store) SaveEvent(e model.Event) error { return put(s, buckets["events"], e.ID, e) }
func (s *Store) GetEvent(id string) (model.Event, error) {
	return get[model.Event](s, buckets["events"], id)
}
func (s *Store) SaveAudit(a model.Audit) error { return put(s, buckets["audits"], a.ID, a) }
func (s *Store) ListAudits(entity string) ([]model.Audit, error) {
	var out []model.Audit
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["audits"]).ForEach(func(_, v []byte) error {
			var a model.Audit
			if e := json.Unmarshal(v, &a); e != nil {
				return e
			}
			if entity == "" || a.Entity == entity {
				out = append(out, a)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket")
		}
		return b.ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
func (s *Store) TouchAudit(entity, action, detail string) error {
	return s.SaveAudit(model.NewAudit(entity+time.Now().Format("150405.000000"), entity, action, detail))
}
