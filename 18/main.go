package main

import (
	"github.com/ValGoldun/study/18/cache"
	v1 "github.com/ValGoldun/study/18/cache/v1"
	"log"
	"time"
)

func main() {
	m()

	cache := cache.New[string, string]()

	go read(cache)
	go write(cache)
	time.Sleep(10 * time.Second)
}

func read(c *cache.Cache[string, string]) {
	for {
		c.Get("")
	}
}

func write(c *cache.Cache[string, string]) {
	for {
		c.Set("", "", time.Second)
	}
}

func m() {
	cache := v1.New()

	cache.Set("data", "secret", time.Second*5)
	log.Println(cache.Get("data")) //проверяем что значение есть
	time.Sleep(time.Second * 6)
	log.Println(cache.Get("data")) //проверяем что значения нет
}
