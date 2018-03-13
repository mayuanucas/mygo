package problem39

import (
	"testing"
	"fmt"
)

func Test_moreThanHalfNum(t *testing.T) {
	numbers := []int{0, 0, 0, 1, 1, 1, 2, 3, 1, 1, 1}

	fmt.Println(moreThanHalfNum(numbers))
}
