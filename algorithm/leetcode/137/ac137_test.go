package problem137

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums1 := []int{2, 2, 3, 2}

	fmt.Println(singleNumber(nums1))
	fmt.Println(singleNumber4(nums1))
}
