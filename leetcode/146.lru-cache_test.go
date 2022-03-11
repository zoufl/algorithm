package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	data     map[int]int
	expired  []int
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		data:     make(map[int]int),
		expired:  make([]int, 0),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.data[key]; ok {
		c.UpdateExpired(key)

		return v
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.data[key]; ok {
		c.data[key] = value
		c.UpdateExpired(key)

		return
	}

	if len(c.data) == c.capacity {
		item := c.expired[0]
		c.expired = c.expired[1:]
		delete(c.data, item)
	}

	c.data[key] = value
	c.UpdateExpired(key)
}

func (c *LRUCache) UpdateExpired(key int) {
	for i, v := range c.expired {
		if v == key {
			c.expired = append(c.expired[:i], c.expired[i+1:]...)
		}
	}

	c.expired = append(c.expired, key)
}

func TestLRUCache(t *testing.T) {
	lRUCache := NewLRUCache(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	t.Log(lRUCache)
	t.Log(lRUCache.Get(1))
	t.Log(lRUCache)
	lRUCache.Put(3, 3)
	t.Log(lRUCache)
	t.Log(lRUCache.Get(2))
	lRUCache.Put(4, 4)
	t.Log(lRUCache)
	t.Log(lRUCache.Get(1))
	t.Log(lRUCache.Get(3))
	t.Log(lRUCache.Get(4))
	t.Log(lRUCache)
}
