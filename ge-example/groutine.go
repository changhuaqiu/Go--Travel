package main

import (
	"fmt"
	"time"
)

func f(name string) {
	fmt.Println(name)
}

func main() {
	go f("kk")
	time.Sleep(time.Second)
}
