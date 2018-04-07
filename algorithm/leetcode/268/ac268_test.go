package problem268

import (
	"testing"
	"fmt"
)

func Test_missingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	nums2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber(nums2))
}
