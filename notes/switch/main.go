package main

import (
	"fmt"
	"runtime"
	"time"
)

// switch 的 case 语句从上到下顺序执行，直到匹配成功时停止。
// switch 只运行选定的 case, 除非以 fallthrough 结束，否则分支自动终止。
// switch 的 case 无需为常数，且取值不必为整数。
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	test(3.1415926)
	test(1)
	test(1.0)
	test(2.71828)

	test2()




}

func test(x float64) {
	switch x {
	case 3.1415926:
		fmt.Println("PI")
	case 1.0:
		fmt.Println("1 or 1.0")
	default:
		fmt.Println("no match", x)
	}
}

// 没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。
func test2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
