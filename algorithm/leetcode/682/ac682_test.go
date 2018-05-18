package problem682

import (
	"testing"
	"fmt"
)

func Test_calPoints(t *testing.T) {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}
