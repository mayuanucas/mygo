package problem57

import (
	"testing"
	"fmt"
)

func Test_findNumbers(t *testing.T) {
	data := []int{1, 2, 4, 7, 11, 15}
	sum := 15

	fmt.Println(findNumbers(data, sum))

	findContinuousSequence(15)
}
