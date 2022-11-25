package cache

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	m map[K]entity[V]
	sync.RWMutex
}

type entity[V any] struct {
	value    V
	deadline time.Time
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		m: make(map[K]entity[V]),
	}
}

func (c *Cache[K, V]) Set(key K, value V, lifetime time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.m[key] = entity[V]{value: value, deadline: time.Now().Add(lifetime)}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.RLock()
	defer c.RUnlock()

	e, exist := c.m[key]
	if !exist {
		return e.value, false
	}

	if time.Now().After(e.deadline) {
		return e.value, exist
	}

	return e.value, true
}
