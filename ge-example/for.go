package main

import (
	"fmt"
)

func main() {
	i := 1
	//第一种遍历
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//第二种遍历
	for j := 1; j <= 9; j++ {
		fmt.Println(j)
	}

}
