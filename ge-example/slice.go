package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	//除了数组所具有的基本操作外，slice还支持比数组更丰富的操作。
	//对切片操作返回的原数据地址上的引用
	a := append(s, "d")
	a = append(a, "cc")
	a = append(a, "dd")
	fmt.Println("a:", a)

	//slice复制
	c := make([]string, len(a))
	copy(c, a)
	fmt.Println("c:", c)

	//slick操作
	l := a[2:5]
	fmt.Println("l:", l)
	l[2] = "43"
	fmt.Println("a:", a)

	//
	d := a[:]
	fmt.Println("d:", d)

	//slice是可以自动增长的
	fmt.Println(len(d), cap(d))

	//防止切片陷阱 使用copy
	//切片返回的是同一片内存区域
}
