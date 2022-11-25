package main

import (
	"github.com/ValGoldun/study/18/cache"
	"log"
	"time"
)

func main() {
	cache := cache.New[string, string]()

	cache.Set("data", "secret", time.Second*5)
	log.Println(cache.Get("data"))

	time.Sleep(time.Second * 6)
	log.Println(cache.Get("data"))
}
