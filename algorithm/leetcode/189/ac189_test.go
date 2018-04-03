package problem189

import (
	"testing"
	"fmt"
)

func Test_rotate(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(numbers, 3)

	fmt.Println(numbers)

	numbers2 := []int{-1}
	rotate(numbers2, 2)
	fmt.Println(numbers2)

}
