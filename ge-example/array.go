package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	//获取数组值的方法跟传统方法类似
	a[4] = 100

	//初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)
	var towD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("2d:", towD)
	fmt.Println(len(b))
}
