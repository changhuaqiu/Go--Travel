package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2009, 11, 17, 20, 34, 58, 2121232, time.UTC)
	p(then)
}
