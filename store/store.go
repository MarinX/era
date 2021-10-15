package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"reflect"

	"github.com/boltdb/bolt"
)

type ErrNotFound struct {
	id string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.id)
}

var (
	ErrSlicePtrNeeded = errors.New("provided target must be a pointer to slice")
)

type Store struct {
	db       *bolt.DB
	location string
}

func New(dir string, buckets ...string) (*Store, error) {
	location := path.Join(dir, "data.db")
	db, err := bolt.Open(location, 0666, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(t *bolt.Tx) error {
		for _, b := range buckets {
			_, err := t.CreateBucketIfNotExists([]byte(b))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return &Store{
		db:       db,
		location: location,
	}, err
}

func (s *Store) Create(key, id string, value interface{}) error {
	return s.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(key))
		buff, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), buff)
	})
}

func (s *Store) ReadAll(key string, value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Ptr || reflect.Indirect(ref).Kind() != reflect.Slice {
		return ErrSlicePtrNeeded
	}

	sliceType := reflect.Indirect(ref).Type().Elem()
	results := reflect.MakeSlice(reflect.Indirect(ref).Type(), 0, 0)
	err := s.db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(key))
		return b.ForEach(func(k, v []byte) error {
			model := reflect.New(sliceType.Elem()).Interface()
			err := json.Unmarshal(v, model)
			if err != nil {
				return err
			}
			results = reflect.Append(results, reflect.ValueOf(model))
			return nil
		})
	})
	reflect.Indirect(ref).Set(results)

	return err
}

func (s *Store) Read(key, id string, value interface{}) error {
	return s.db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(key))
		v := b.Get([]byte(id))
		if v == nil {
			return &ErrNotFound{id: id}
		}
		return json.Unmarshal(v, value)
	})
}

func (s *Store) Update(key, id string, value interface{}) error {
	return s.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(key))
		buff, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), buff)
	})
}

func (s *Store) Delete(key, id string) error {
	return s.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(key))
		return b.Delete([]byte(id))
	})
}

func (s *Store) Close() error {
	return s.db.Close()
}
