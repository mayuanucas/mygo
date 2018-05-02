package problem575

import (
	"testing"
	"fmt"
)

func Test_distributeCandies(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}

	fmt.Println(distributeCandies(nums))

	nums2 := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(nums2))
}
