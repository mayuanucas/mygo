package problem438

import (
	"testing"
	"fmt"
)

func Test_findAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, p))

	s1 := "abab"
	p1 := "ab"
	fmt.Println(findAnagrams(s1, p1))

	s2 := "ababababab"
	p2 := "aab"
	fmt.Println(findAnagrams(s2, p2))
}
