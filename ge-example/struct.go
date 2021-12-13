package main

import "fmt"

type Person struct {
	name string
	age  int64
}

func main() {
	fmt.Println(Person{"a", 1})
	fmt.Println(Person{name: "a", age: 1})
	a := &Person{"a", 1}
	fmt.Println(a.name)
}
