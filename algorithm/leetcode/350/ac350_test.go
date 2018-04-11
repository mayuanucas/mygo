package problem350

import (
	"testing"
	"fmt"
)

func Test_intersect(t *testing.T) {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}

	fmt.Println(intersect(num1, num2))
}
