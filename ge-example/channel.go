package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	f := func() {
		message <- "ping"
		fmt.Println("message getted")
	}
	go f()
	fmt.Println("begin to get message")
	time.Sleep(time.Second * 5)
	msg := <-message
	fmt.Println(msg)
	time.Sleep(time.Second * 1)
}
