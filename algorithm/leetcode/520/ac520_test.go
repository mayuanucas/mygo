package problem520

import (
	"testing"
	"fmt"
)

func Test_detectCapitalUse(t *testing.T) {
	str1 := "USA"
	fmt.Println(detectCapitalUse(str1))

	str2 := "Leetcode"
	fmt.Println(detectCapitalUse(str2))

	str3 := "LEETcode"
	fmt.Println(detectCapitalUse(str3))

	fmt.Println(detectCapitalUse2(str1))
	fmt.Println(detectCapitalUse2(str2))
	fmt.Println(detectCapitalUse2(str3))
}
