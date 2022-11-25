package main

import (
	"fmt"
	"time"
)

func main() {
	h := &human{
		FullName:  "Петр",
		Sex:       "Мужчина",
		BirthDate: time.Date(1989, 11, 0, 0, 0, 0, 0, time.UTC),
		Height:    176,
		Weight:    80,
	}

	fmt.Println(h.getFullName())
	fmt.Println(h.getSex())
	fmt.Println(h.getAge())
	fmt.Println(h.getHeight())
	fmt.Println(h.getWeight())
}

type human struct {
	FullName  string
	Sex       string
	BirthDate time.Time
	Height    uint16
	Weight    uint16
}

func (h *human) getFullName() string {
	return h.FullName
}

func (h *human) getSex() string {
	return h.Sex
}

func (h *human) getAge() uint8 {
	return uint8(time.Now().Sub(h.BirthDate).Hours() / (24 * 365))
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

func (h *human) setBirthDate(v time.Time) {
	h.BirthDate = v
}

func (h *human) setHeight(v uint16) {
	h.Height = v
}

func (h *human) setWeight(v uint16) {
	h.Weight = v
}
