package main

import (
	"fmt"
)

func main() {
	// nil 是不能比较的， == 对于 nil 来说是一种未定义的操作。
	// 运行下面这行将报错: invalid operation: nil == nil (operator == not defined on nil)
	// fmt.Println(nil == nil)

	// nil没有type
	fmt.Printf("%T\n", nil)
	fmt.Println("**************************")

	// 不同类型 nil 的 address 是一样的
	var number *int
	var dict map[string]int
	fmt.Printf("%p\n", number)
	fmt.Printf("%p\n", dict)
	fmt.Println("**************************")

	// nil 是 map，slice，pointer，channel，func，interface 的零值
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)
	fmt.Println("**************************")

}
