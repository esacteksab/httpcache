// Package boltdbcache provides an implementation of httpcache.ContextCache that
// uses go.etcd.io/bbolt
package bboltdbcache

import (
	"context"

	"go.etcd.io/bbolt"
)

// Cache is an implementation of httpcache.Cache with boltdb storage
type Cache struct {
	db     *bbolt.DB
	bucket []byte
}

// Get returns the response corresponding to key if present
func (c *Cache) Get(_ context.Context, key string) (resp []byte, ok bool, err error) {
	if err := c.db.View(func(tx *bbolt.Tx) error {
		tmp := tx.Bucket(c.bucket).Get([]byte(key))
		if tmp != nil {
			resp = make([]byte, len(tmp))
			copy(resp, tmp)
			ok = true
		}
		return nil
	}); err != nil {
		return nil, false, err
	}
	return
}

// Set saves a response to the cache as key
func (c *Cache) Set(_ context.Context, key string, resp []byte) error {
	return c.db.Batch(func(tx *bbolt.Tx) error {
		return tx.Bucket(c.bucket).Put([]byte(key), resp)
	})
}

// Delete removes the response with key from the cache
func (c *Cache) Delete(_ context.Context, key string) error {
	return c.db.Batch(func(tx *bbolt.Tx) error {
		return tx.Bucket(c.bucket).Delete([]byte(key))
	})
}

// New returns a new Cache that will use the provided boltdb as underlying storage
func New(db *bbolt.DB, bucket []byte) *Cache {
	return &Cache{db: db, bucket: bucket}
}

// Open opens a boltdb at path and returns a Cache using the provided bucket
func Open(path string, bucket []byte) (*bbolt.DB, *Cache, error) {
	if bucket == nil {
		bucket = []byte("cache")
	}
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, nil, err
	}
	if err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	}); err != nil {
		return nil, nil, err
	}
	return db, New(db, bucket), nil
}
