package problem349

import (
	"testing"
	"fmt"
)

func Test_intersection(t *testing.T) {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}

	fmt.Println(intersection(num1, num2))
}
