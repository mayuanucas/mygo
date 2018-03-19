package problem56

import (
	"testing"
	"fmt"
)

func Test_findNumbersAppearOnce(t *testing.T) {
	data := []int{2, 4, 3, 6, 3, 2, 5, 5}
	fmt.Println(findNumbersAppearOnce(data))

	data2 := []int{2, 2, 2, 3, 3, 3, 4}
	fmt.Println(findNumberAppearOnce2(data2))
}
