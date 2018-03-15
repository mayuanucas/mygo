package problem47

import (
	"testing"
	"fmt"
)

func Test_getMaxValue(t *testing.T) {
	matrix := [][]int{
		{1, 3, 3},
		{2, 1, 3},
		{2, 2, 1},
	}
	fmt.Println(getMaxValue(matrix))
	fmt.Println(getMaxValue2(matrix))
}
