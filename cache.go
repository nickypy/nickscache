package main

import (
	"sync"
)

// Cache object definition
type Cache struct {
	Cache   map[string]*Node
	MaxSize int
	Order   *List
	mutex   *sync.Mutex
}

// NewCache return a newly constructed cache
func NewCache(maxSize int) *Cache {
	return &Cache{make(map[string]*Node), maxSize, NewList(), new(sync.Mutex)}
}

// Put adds a string to the cache if possible, otherwise evict from the cache then insert
func (cache *Cache) Put(key, value string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// evict from the cache
	if cache.Order.Length == cache.MaxSize {
		toRemove := cache.Order.Head
		cache.Order.RemoveNode(toRemove)

		delete(cache.Cache, toRemove.MapKey)
	}

	if n, exists := cache.Cache[key]; exists {
		cache.Order.RemoveNode(n)
	}

	node := cache.Order.Insert(value, key)
	cache.Cache[key] = node
}

// Get a key from the map, if it exists otherwise, return an empty string
func (cache *Cache) Get(key string) NodeKey {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if value, exists := cache.Cache[key]; exists {
		removed := cache.Order.RemoveNode(value)
		cache.Order.Insert(removed.Key, key)

		return removed.Key
	}

	return ""
}

// SizeRemaining return space left in bytes
func (cache Cache) SizeRemaining() int {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	return cache.MaxSize - cache.Order.Length
}
