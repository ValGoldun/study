package cache

import "time"

func New[K comparable, V any]() *cache[K, V] {
	return &cache[K, V]{
		m: make(map[K]entity[V]),
	}
}

func (c *cache[K, V]) Set(key K, value V, lifetime time.Duration) {
	c.m[key] = entity[V]{value: value, deadline: time.Now().Add(lifetime)}
}

func (c *cache[K, V]) Get(key K) (V, bool) {
	entity, exist := c.m[key]
	if !exist {
		return entity.value, false
	}

	if time.Now().After(entity.deadline) {
		return entity.value, false
	}

	return entity.value, true
}

type cache[K comparable, V any] struct {
	m map[K]entity[V]
}

type entity[V any] struct {
	value    V
	deadline time.Time
}
