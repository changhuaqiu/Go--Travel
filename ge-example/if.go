package main

import "fmt"

func main() {
	//go中的if-else的用法跟c++中的用法类似。圆括号不是必须的
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//可以在if 语句中直接赋值
	if n := 1; n > 0 {
		fmt.Println("n > 0")
	}
}
