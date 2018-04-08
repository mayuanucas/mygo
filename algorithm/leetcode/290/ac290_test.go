package problem290

import (
	"testing"
	"fmt"
)

func Test_wordPattern(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat fish"

	fmt.Println(wordPattern(pattern, str))
}
