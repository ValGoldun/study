package v1

import (
	"time"
)

type Cache struct {
	m map[string]entity
}

type entity struct {
	value    string
	deadline time.Time
}

func New() *Cache {
	return &Cache{
		m: make(map[string]entity),
	}
}

func (c *Cache) Set(key string, value string, lifetime time.Duration) {
	c.m[key] = entity{value: value, deadline: time.Now().Add(lifetime)}
}

func (c *Cache) Get(key string) (string, bool) {
	e, exist := c.m[key]
	if !exist {
		return e.value, false
	}

	if time.Now().After(e.deadline) {
		return e.value, false
	}

	return e.value, true
}
