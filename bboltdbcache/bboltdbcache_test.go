package bboltdbcache

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestBoltDBCache(t *testing.T) {
	ctx := context.Background()

	tempDir, err := os.MkdirTemp("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	db, cache, err := Open(filepath.Join(tempDir, "db"), nil)
	if err != nil {
		t.Fatalf("New leveldb,: %v", err)
	}
	defer db.Close()

	key := "testKey"
	if _, ok, err := cache.Get(ctx, key); err != nil {
		t.Fatalf("cache.Get failed: %v", err)
	} else if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	if err := cache.Set(ctx, key, val); err != nil {
		t.Fatalf("cache.Set failed: %v", err)
	}

	if retVal, ok, err := cache.Get(ctx, key); err != nil {
		t.Fatalf("cache.Get failed: %v", err)
	} else if !ok {
		t.Fatal("could not retrieve an element we just added")
	} else if !bytes.Equal(retVal, val) {
		log.Println(retVal, val)
		t.Fatal("retrieved a different value than what we put in")
	}

	if err := cache.Delete(ctx, key); err != nil {
		t.Fatalf("cache.Delete failed: %v", err)
	}

	if _, ok, err := cache.Get(ctx, key); err != nil {
		t.Fatalf("cache.Get failed: %v", err)
	} else if ok {
		t.Fatal("deleted key still present")
	}
}
