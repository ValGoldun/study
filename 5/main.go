package main

import "fmt"

func main() {
	var array []int

	for i := 0; i < 10; i++ {
		array = append(array, i)
	}

	for _, i := range array {
		fmt.Println(i)
	}
}
