package main

import "fmt"

type Animal interface {
	Speack() string
}

type Cat struct {
	name string
}

func (c Cat) Speack() string {
	return "Cat"
}
func PrintAll(a Animal) {
	fmt.Println(a.Speack())
}
func main() {
	c := new(Cat)
	PrintAll(c)

}
