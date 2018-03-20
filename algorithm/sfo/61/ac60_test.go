package problem61

import (
	"testing"
	"fmt"
)

func Test_isContinuous(t *testing.T) {
	data := []int{3, 4, 0, 5, 1}
	fmt.Println(isContinuous(data))
}
