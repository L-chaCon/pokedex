package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entrys map[string]cacheEntry
	mux    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entrys: make(map[string]cacheEntry),
		mux:    &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entrys[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cache, ok := c.entrys[key]
	return cache.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.entrys {
		if time.Since(entry.createdAt) > interval {
			delete(c.entrys, key)
		}
	}
}
