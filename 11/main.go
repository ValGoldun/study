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

	fmt.Println(h.getFullName())
	fmt.Println(h.getSex())
	fmt.Println(h.getAge())
	fmt.Println(h.getHeight())
	fmt.Println(h.getWeight())
}

type human struct {
	FullName string
	Sex      string
	Age      uint8
	Height   uint16
	Weight   uint16
}

func (h human) getFullName() string {
	return h.FullName
}

func (h human) getSex() string {
	return h.Sex
}

func (h human) getAge() uint8 {
	return h.Age
}

func (h human) getHeight() uint16 {
	return h.Height
}

func (h human) getWeight() uint16 {
	return h.Weight
}
