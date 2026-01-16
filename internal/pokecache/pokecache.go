package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mutex:        sync.Mutex{},
		interval:     interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			age := now.Sub(entry.createdAt)
			if age > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}
