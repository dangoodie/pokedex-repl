package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data      []byte
	timestamp time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.cache[key] = cacheEntry{
		data:      data,
		timestamp: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	if time.Since(entry.timestamp) > time.Minute {
		return nil, false
	}

	return entry.data, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		c.reap(interval)
		time.Sleep(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	for key, entry := range c.cache {
		if time.Since(entry.timestamp) > interval {
			delete(c.cache, key)
		}
	}
}