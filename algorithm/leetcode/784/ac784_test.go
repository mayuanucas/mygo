package problem784

import (
	"testing"
	"fmt"
)

func Test_letterCasePermutation(t *testing.T) {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))

	s1 := ""
	fmt.Println(letterCasePermutation(s1))
}
