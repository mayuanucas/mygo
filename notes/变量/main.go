package main

import (
	"fmt"
	"math"
)

func main() {
	// 基本类型
	var (
		a bool
		b string
		c byte
		d int
		e rune
		f float64
		g complex64
	)

	fmt.Printf("%v %q %v %v %v %v %v", a, b, c, d, e, f, g)

	// 变量初始化，变量声明可以包含初始化值。如果初始化值已存在，可以省略类型，变量从初始值获得类型。
	var i, j, k = 12.99, true, "test"
	fmt.Println(i, j, k)
	fmt.Printf("i: %T, j:%T, k: %T\n", i, j, k)

	// 在函数中,简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
	m := 3
	n := false
	str2 := "golang"
	fmt.Println(m, n, str2)

	// 类型转换
	var x, y = 3, 4
	var ans = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("ans: %v, type: %T\n", ans, ans)

}
