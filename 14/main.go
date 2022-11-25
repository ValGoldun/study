package main

import "fmt"

func main() {
	cdek := CDEK{}
	russianPost := RussianPost{}

	Delivery(cdek)
	Delivery(russianPost)
}

func Delivery(d Deliverier) {
	d.Delivery()
}

type Deliverier interface {
	Delivery()
}

type CDEK struct{}

func (c CDEK) Delivery() {
	fmt.Println("Доставлено СДЕК")
}

type RussianPost struct{}

func (r RussianPost) Delivery() {
	fmt.Println("Доставлено ПОЧТА")
}
