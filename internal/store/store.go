package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"

	"github.com/boltdb/bolt"
)

type ErrNotFound struct {
	id string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.id)
}

type Store struct {
	db       *bolt.DB
	location string
}

func New(dir string) (*Store, error) {
	location := path.Join(dir, "data.db")
	db, err := bolt.Open(location, 0666, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte("contacts"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("keys"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("settings"))
		return err
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
	return s.db.View(func(t *bolt.Tx) error {
		//@TODO
		return errors.New("//TODO")
	})
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
