// Package badgerdbcache provides an implementation of httpcache.Cache that
// uses github.com/dgraph-io/badger
package badgerdbcache

import (
	"context"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

// Cache is an implementation of httpcache.Cache with boltdb storage
type Cache struct {
	db *badger.DB
}

// Get returns the response corresponding to key if present
func (c *Cache) Get(_ context.Context, key string) (resp []byte, ok bool, err error) {
	err = c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}
		resp, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	if resp != nil && err == nil {
		ok = true
	}
	return
}

// Set saves a response to the cache as key
func (c *Cache) Set(_ context.Context, key string, resp []byte) error {
	log.Println("Set", key, string(resp))
	return c.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), resp)
	})
}

// Delete removes the response with key from the cache
func (c *Cache) Delete(_ context.Context, key string) error {
	return c.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// New returns a new Cache that will store a boltdb in path using bucket (defaults to "cache")
func New(db *badger.DB) *Cache {
	return &Cache{db: db}
}
