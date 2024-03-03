package pokecache

import (
    "sync"
    "time"
)

type Cache struct {
    entries   map[string]cacheEntry
    mu        *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        mu:      &sync.Mutex{},
        entries: make(map[string]cacheEntry),
    }

    go cache.reapLoop(interval)

    return cache
}

func (c *Cache) Add(key string, value []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.entries[key] = cacheEntry{
        createdAt: time.Now(),
        val: value,
    }
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    value, ok := c.entries[key]
    return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    for key, entry := range c.entries {
        if entry.createdAt.Before(now.Add(-last)) {
            delete(c.entries, key)
        }
    }
}
