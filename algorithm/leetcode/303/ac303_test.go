package problem303

import (
	"testing"
	"fmt"
)

func TestNumArray_SumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}

	obj := Constructor(nums)
	fmt.Println(obj.data)
	fmt.Println(obj.SumRange(0, 2))
}
