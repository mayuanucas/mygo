package problem581

import (
	"testing"
	"fmt"
)

func Test_findUnsortedSubarray(t *testing.T) {
	num := []int{2, 6, 4, 8, 10, 9, 15}
	num2 := []int{1, 2, 3, 4}

	fmt.Println(findUnsortedSubarray(num))
	fmt.Println(findUnsortedSubarray(num2))
}
