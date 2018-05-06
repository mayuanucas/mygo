package problem830

import (
	"testing"
	"fmt"
)

func Test_largeGroupPositions(t *testing.T) {
	str1 := "abbxxxxzzy"
	fmt.Println(largeGroupPositions(str1))

	str2 := "abc"
	fmt.Println(largeGroupPositions(str2))

	str3 := "abcdddeeeeaabbbcd"
	fmt.Println(largeGroupPositions(str3))

	str4 := "aaa"
	fmt.Println(largeGroupPositions(str4))
}
