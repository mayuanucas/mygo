package problem53

import (
	"testing"
	"fmt"
)

func Test_getNumberOfK(t *testing.T) {
	numbers := []int{1, 2, 3, 3, 3, 3, 4, 5}
	k := 3
	fmt.Println(getNumberOfK(numbers, k))
}
