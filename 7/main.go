package main

import "fmt"

func main() {
	var a = 20
	var b = 4

	calculate(a, b, add)
	calculate(a, b, sub)
	calculate(a, b, mul)
	calculate(a, b, div)
}

func calculate(a int, b int, calculator func(a int, b int) int) {
	fmt.Println(calculator(a, b))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}
