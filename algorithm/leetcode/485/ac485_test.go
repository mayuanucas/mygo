package problem485

import (
	"testing"
	"fmt"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Println(findMaxConsecutiveOnes(nums))
}
