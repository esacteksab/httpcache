package test_test

import (
	"testing"

	"github.com/esacteksab/httpcache"
	"github.com/esacteksab/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
