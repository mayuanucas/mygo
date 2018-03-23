package problem035

import (
	"testing"
	"fmt"
)

func Test_searchInsert(t *testing.T) {
	data := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(data, target))
}
