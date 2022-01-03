package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//匿名函数
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func GET(i func(p string, q string) string) {
	fmt.Println(i("GET", "for"))
}
func main() {
	sum(1, 2)
	sum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//匿名函数
	//函数变量带有状态
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	//	普通类型
	var f = func() {
		fmt.Println("hello world")
	}
	f()
	//匿名函数作为参数
	var f1 = func(p string, q string) string {
		return p + q + "ret"
	}
	GET(f1)

}
