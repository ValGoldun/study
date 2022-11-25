package main

import "fmt"

func main() {
	h := human{
		FullName: "Петр",
		Sex:      "Мужчина",
		Age:      33,
		Height:   176,
		Weight:   80,
	}

	fmt.Printf(
		"Имя: %s\nПол: %s\nВозраст: %d\nРост: %d\nВес: %d\n",
		h.FullName,
		h.Sex,
		h.Age,
		h.Height,
		h.Weight,
	)
}

type human struct {
	FullName string
	Sex      string
	Age      uint8
	Height   uint16
	Weight   uint16
}
