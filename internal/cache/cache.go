package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// cachePtr ... pointer to in-memory cache
var cachePtr *cache.Cache

var objOnce sync.Once

const (
	// For use with cache values which are not going to change once app is up.
	NoExpiration = cache.NoExpiration

	// For Setting default value of expiration, i.e. NoExpiration
	DefaultExpiration = cache.DefaultExpiration
)

// GetCacheInstance returns pointer to cache singleton
func GetCacheInstance() *cache.Cache {
	if cachePtr != nil {
		return cachePtr
	}

	objOnce.Do(func() {
		cachePtr = cache.New(cache.NoExpiration, cache.NoExpiration)

	})
	return cachePtr
}

// Get returns value in cache for given key. If key is not found, it returns nil.
func Get(key string) interface{} {
	cachePtr := GetCacheInstance()
	if value, found := cachePtr.Get(key); found {
		return value
	}
	return nil
}

// Set sets the given value in cache for given TTL
func Set(key string, value interface{}, ttl time.Duration) {
	cachePtr := GetCacheInstance()
	cachePtr.Set(key, value, ttl)
}