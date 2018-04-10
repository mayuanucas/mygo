package problem345

import (
	"testing"
	"fmt"
)

func Test_reverseVowels(t *testing.T) {
	str := "hello"
	str2 := "leetcode"
	fmt.Println(reverseVowels(str))
	fmt.Println(reverseVowels(str2))
}
