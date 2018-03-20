package problem59

import (
	"testing"
	"fmt"
)

func Test_maxInWindows(t *testing.T) {
	data := []int{2, 3, 4, 2, 6, 2, 5, 1}
	size := 3

	fmt.Println(maxInWindows(data, size))
}
