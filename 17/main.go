package main

import "sync"

func main() {
	var m = make(map[int]int)
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func(mutex *sync.Mutex) {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}(&mutex)
	}
}
