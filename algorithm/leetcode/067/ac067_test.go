package problem067

import (
	"testing"
	"fmt"
)

func Test_addBinary(t *testing.T) {
	a := "11"
	b := "01"
	fmt.Println(addBinary(a, b))
	fmt.Println(addBinary2(a, b))
}
