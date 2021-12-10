package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["ke1"] = 2
	m["ke2"] = 1
	fmt.Println("m:", m)
	if val, ok := m["ke1"]; ok {
		fmt.Println("val:", val)
	}
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
