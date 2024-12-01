package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val         []byte
	createdTime time.Time
}

func newCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:         val,
		createdTime: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cacheExpirationTime := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdTime.Before(cacheExpirationTime) {
			delete(c.cache, k)
		}
	}

}
