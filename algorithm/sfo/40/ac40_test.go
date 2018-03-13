package problem40

import (
	"testing"
	"fmt"
)

func Test_getLeastNumbers(t *testing.T) {
	numbers := []int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ans := getLeastNumbers(numbers, 3)
	for _, value := range ans {
		fmt.Println(value)
	}

}
