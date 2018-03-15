package problem45

import (
	"testing"
	"fmt"
)

func Test_printMinNumber(t *testing.T) {
	numbers := []int{3, 32, 321}
	fmt.Println(printMinNumber(numbers))
}
