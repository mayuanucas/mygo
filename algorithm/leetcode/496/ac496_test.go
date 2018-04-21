package problem496

import (
	"testing"
	"fmt"
)

func Test_nextGreaterElement(t *testing.T) {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(num1, num2))
}
