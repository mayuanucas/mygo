package problem438

import (
	"testing"
	"fmt"
)

func Test_findAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, p))
}
