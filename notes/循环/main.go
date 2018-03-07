package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)

	// for 是 Go 中的 “while”
	// 此时可以去掉分号，因为 C 的 while 在 Go 中叫 for
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	// 省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑
	// for {
	//
	// }

}
