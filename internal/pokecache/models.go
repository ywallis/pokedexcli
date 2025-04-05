package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Map map[string]CacheEntry
	mu  sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[key] = CacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.Map[key]

	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(lifetime time.Duration) {
	ticker := time.NewTicker(lifetime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			c.mu.Lock()
			for key, val := range c.Map {
				age := time.Since(val.createdAt)
				if age >= lifetime {
					delete(c.Map, key)

				}
			}
			c.mu.Unlock()
		}
	}
}

func NewCache(lifetime time.Duration) *Cache {

	Cache := Cache{make(map[string]CacheEntry), sync.Mutex{}}

	go Cache.reapLoop(lifetime)

	return &Cache
}
