package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type Cache struct {
    CachedResults map[string]cacheEntry
    mu *sync.Mutex
}

func NewCache(interval time.Duration) Cache{
    cache := Cache { CachedResults: map[string]cacheEntry{}, 
    mu: &sync.Mutex{}, 
    }
    go cache.reapLoop(interval)
    return cache
}

 func (c *Cache) Add(key string, value[]byte){
    c.mu.Lock()
    defer c.mu.Unlock()
    _, exists := c.CachedResults[key]
    if !exists {
        c.CachedResults[key] = cacheEntry{ val: value, createdAt: time.Now() }
    }

}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    entry, exists := c.CachedResults[key]
    if exists {
        return entry.val, true
    }
    return nil, false
}

func (c * Cache) reapLoop(interval time.Duration) {
    

    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.mu.Lock()
        for key, entry := range c.CachedResults {
            if time.Since(entry.createdAt) > interval {
                delete(c.CachedResults, key)
            }
        }
        c.mu.Unlock()
    }
}