package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (people Person) Getname() {
	people.name = "aa"
}

func (people Person) Getage() int {
	return people.age
}

func main() {
	people := Person{"kk", 25}
	people.Getname()
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.Getage())
	//上面是变量 存在一个自动变量转换的问题
	u := &people
	u.Getage()

}
