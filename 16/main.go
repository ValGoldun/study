package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
