package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache 		map[string]cacheEntry
	mux			*sync.Mutex
}
type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
} 

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache:	make(map[string]cacheEntry),
		mux:	&sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add (key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt:	time.Now().UTC(),
		val:		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for key, val := range c.cache {
			if val.createdAt.Before(time.Now().UTC().Add(-interval)) {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}