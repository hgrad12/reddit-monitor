package internal

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewRedditCache()
	cache.cache.Set("key", Result{ Name: "name"}, time.Microsecond)

	t.Run("retrieve an item that doesn't exist", func(t *testing.T) {
		t.Parallel()
		key := "key1"

		res := cache.Get(key)

		if res.Name != "" {
			t.Errorf("the cache retrieved a value %+v from key %s", res, key)
		}
	})

	t.Run("retrieve an item that exists", func(t *testing.T) {
		t.Parallel()
		
		key := "key"
		res := cache.Get(key)

		if res.Name == "" {
			t.Errorf("the cache retrieved a value %+v from key %s", res, key)
		}
	})
}