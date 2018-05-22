package problem686

import (
	"testing"
	"fmt"
)

func Test_repeatedStringMatch(t *testing.T) {
	a := "abababaaba"
	b := "aabaaba"

	fmt.Println(repeatedStringMatch(a, b))

	aa := "bb"
	bb := "bbbbbbb"
	fmt.Println(repeatedStringMatch(aa, bb))
}
