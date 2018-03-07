package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2, 3, 99))
	fmt.Println(myPow(2, 3, 5))
	fmt.Printf("sqrt %v is %v", 2, Sqrt(2))
}

// if 语句可以在表达式前执行一个简单的语句，该语句声明的变量作用域仅在 if 之内。
// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
func myPow(x, y, ans float64) float64 {
	if z := math.Pow(x, y); z < ans {
		return z
	} else {
		fmt.Printf("z type: %T\n", z)
	}
	return ans
}

// 牛顿法实现平方根函数
func Sqrt(x float64) float64 {
	if 0 >= x {
		return 0
	}

	z := 1.0
	for {
		temp := z - (z*z-x)/(2*z)
		if math.Abs(temp-z) < 1e-12 {
			break
		}
		z = temp
	}
	return z
}
