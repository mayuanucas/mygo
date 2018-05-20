package problem665

import (
	"testing"
	"fmt"
)

func Test_checkPossibility(t *testing.T) {
	nums := []int{4, 2, 1}
	fmt.Println(checkPossibility(nums))

	nums2 := []int{4, 2, 3}
	fmt.Println(checkPossibility(nums2))

	nums3 := []int{3, 4, 2, 3}
	fmt.Println(checkPossibility(nums3))
}
