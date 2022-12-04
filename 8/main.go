package main

import "fmt"

func main() {
	var a = []string{"раз", "два", "три"}
	var b = []string{"раз", "два", "три", "четыре"}

	fmt.Println(median(a))
	fmt.Println(median(b))
}

func median(array []string) []string {
	if len(array) < 3 {
		return array
	}

	if len(array)%2 == 0 {
		return array[(len(array)/2)-1 : (len(array)/2)+1]
	}
	return array[(len(array) / 2) : (len(array)/2)+1]
}
