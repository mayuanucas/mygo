package main

import (
	"fmt"
)

// 常量可以是字符、字符串、布尔值或数值
// 常量不能用 := 语法声明
const PI = 3.1415926

func main(){
	fmt.Printf("PI: %v, %T\n", PI, PI)

	const (
		TRUE = true
		FALSE = false
	)
	fmt.Println(TRUE, FALSE)

}