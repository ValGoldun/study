package main

import "fmt"

func main() {
	h := &human{
		FullName: "Петр",
		Sex:      "Мужчина",
		Age:      33,
		Height:   176,
		Weight:   80,
	}

	h.setAge(38)
	h.setWeight(85)

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

func (h *human) getFullName() string {
	return h.FullName
}

func (h *human) getSex() string {
	return h.Sex
}

func (h *human) getAge() uint8 {
	return h.Age
}

func (h *human) getHeight() uint16 {
	return h.Height
}

func (h *human) getWeight() uint16 {
	return h.Weight
}

func (h *human) setFullName(v string) {
	h.FullName = v
}

func (h *human) setSex(v string) {
	h.Sex = v
}

func (h *human) setAge(v uint8) {
	h.Age = v
}

func (h *human) setHeight(v uint16) {
	h.Height = v
}

func (h *human) setWeight(v uint16) {
	h.Weight = v
}
