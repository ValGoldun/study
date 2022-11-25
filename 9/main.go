package main

import "fmt"

func main() {
	var m = make(map[string]string)

	m["Петр"] = "Иванов"
	m["Иван"] = "Петров"

	fmt.Println(m["Петр"])
	delete(m, "Петр")
	fmt.Println(m["Петр"])
}
