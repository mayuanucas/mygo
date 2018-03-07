package main

// import "fmt"
// import "math
import (
	"fmt"
	m "math" // 起别名
	. "time" // 直接使用该包
)

func main() {
	fmt.Printf("Sqrt %v is %v\n", 3, m.Sqrt(3))
	fmt.Println("time is ", Now())
	fmt.Println("PI: ", m.Pi)
}

// 导出
// 在 Go 中，如果一个名字以大写字母开头，那么它就是导出的。例如： Pi 就是导出的，它导出自 math 包。
// 在导入一个包时，只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
