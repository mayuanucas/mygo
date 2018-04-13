package problem414

import (
	"testing"
	"fmt"
)

func Test_thirdMax(t *testing.T) {
	num1 := []int{3, 2, 1}
	num2 := []int{2, 1}
	num3 := []int{2, 3, 2, 1}

	fmt.Println(thirdMax(num1))
	fmt.Println(thirdMax(num2))
	fmt.Println(thirdMax(num3))
}
