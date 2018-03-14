package problem42

import (
	"testing"
	"fmt"
)

func Test_findGreatestSumOfSubArray(t *testing.T) {
	data := []int{1, -2, 3, 10, -4, 7, 2, -5}

	fmt.Println(findGreatestSumOfSubArray(data))

}
