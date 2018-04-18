package problem459

import (
	"testing"
	"fmt"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	s := "ab"
	fmt.Println(repeatedSubstringPattern(s))

	s2 := "abab"
	fmt.Println(repeatedSubstringPattern(s2))

}
