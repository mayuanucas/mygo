package problem643

import (
	"testing"
	"fmt"
)

func Test_findMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}

	fmt.Println(findMaxAverage(nums, 4))

	nums2 := []int{-1}
	fmt.Println(findMaxAverage(nums2, 1))

}
