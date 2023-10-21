package goCacheDemo

import (
	"goCacheDemo/LRU"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *LRU.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = LRU.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (ByteView, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return ByteView{}, false
	}
	if v, isOk := c.lru.Get(key); isOk {
		return v.(ByteView), isOk
	}
	return ByteView{}, false
}
