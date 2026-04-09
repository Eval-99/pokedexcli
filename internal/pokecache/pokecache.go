package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachedItems map[string]cacheEntry
	interval    time.Duration
	mu          *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cachedItems: make(map[string]cacheEntry),
		interval:    interval,
		mu:          &sync.Mutex{},
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	newVal := cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Lock()
	c.cachedItems[key] = newVal
	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	item, ok := c.cachedItems[key]
	if !ok {
		return []byte{}, false
	}
	defer c.mu.Unlock()

	return item.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for name, entry := range c.cachedItems {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cachedItems, name)
			}
		}
		c.mu.Unlock()
	}
}
