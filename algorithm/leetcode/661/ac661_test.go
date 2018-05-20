package problem661

import (
	"testing"
	"fmt"
)

func Test_imageSmoother(t *testing.T) {
	M := [][]int{
		{0, 1, 2, 3},
		{2, 3, 1, 1},
	}

	fmt.Println(imageSmoother(M))
}
