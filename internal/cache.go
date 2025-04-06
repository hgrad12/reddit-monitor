package internal

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get(key string) *Result
	Set(key string, value Result)
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

type RedditCache struct {
	cache *cache.Cache
}

func NewRedditCache() *RedditCache {
	c := cache.New(defaultExpiration, purgeTime)
	return &RedditCache{
		cache: c,
	}
}

func (r *RedditCache) Get(key string) Result {
	val, found := r.cache.Get(key)
	if found {
		return val.(Result)
	}
	return Result{}
}

func (r *RedditCache) Set(key string, value Result) {
	r.cache.Set(key, value, defaultExpiration)
}