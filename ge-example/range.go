package main

import (
	"fmt"
)

func main() {
	//对各种数据结构的遍历
	//slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for idx, _ := range nums {
		fmt.Println("idx:", idx)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s-> %s\n", k, v)
	}
}
