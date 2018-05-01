package problem566

import (
	"testing"
	"fmt"
)

func Test_matrixReshape(t *testing.T) {
	nums := [][]int{
		{1, 2},
		{3, 4}}

	fmt.Println(matrixReshape(nums, 1, 4))
	fmt.Println(matrixReshape(nums, 2, 4))
}
