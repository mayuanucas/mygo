package main

import "fmt"

// 函数可以接受0个或多个参数
func main() {
	fmt.Println("this is main function")
	fmt.Println(Add(9, 12))
	fmt.Println(add(9, 12, 1))
	fmt.Println(add2(9, 12, 1))
	fmt.Println(add3(false, 12, 1))
	fmt.Println(swap("time", "golang"))
	fmt.Println(split(20))
}

func Add(x int, y int) int {
	return x + y
}

func add(x int, y int, z int) int {
	return x + y + z
}

func add2(x, y, z int) int {
	return x + y + z
}

func add3(x bool, y, z int) (bool, int) {
	return x, y + z
}

func swap(x, y string) (string, string) {
	a, b := x, y
	fmt.Print("swap->", a, " ",b, "->")
	return y, x
}

// 命名返回值
// Go 的返回值可以被命名，它们会被视作定义在函数顶部的变量。
// 没有参数的 return 语句返回已命名的返回值，也就是 “直接” 返回。
func split(sum int) (x, y int)  {
	x = sum * 2 - 1
	y = x - sum
	return
}
