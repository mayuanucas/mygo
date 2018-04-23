package problem120

import (
	"testing"
	"fmt"
)

func Test_minimumTotal(t *testing.T) {
	data := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}

	fmt.Println(minimumTotal(data))
}
