package problem594

import (
	"fmt"
	"testing"
)

func Test_findLHS(t *testing.T) {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))

	nums2 := []int{2, 2, 2, 9, 3, 3, 3}
	fmt.Println(findLHS(nums2))
}
