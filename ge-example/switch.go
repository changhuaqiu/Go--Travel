package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	i := 2
	//基本的case语句
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	//在同一个case语句中可以用逗号来分隔多个表达式。
	switch time.Now().Weekday() {
	case time.Saturday, time.Wednesday:
		fmt.Println("It is Week")
	default:
		fmt.Println("It is work day")
	}

	//不带表达式类型的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("< 12")
	default:
		fmt.Println(">12")
	}
	//如果函数的参数是interface{}
	// 通过fmt fmt.Println("%T")
	// 通过reflect reflect.Typeof(v).String()
	//通过类型断言
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("default %T", t)
		}
		fmt.Println(reflect.TypeOf(i).String())
	}
	whatAmI(true)

}
