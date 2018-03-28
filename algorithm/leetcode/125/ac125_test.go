package problem125

import (
	"testing"
	"fmt"
)

func Test_isCharacter(t *testing.T) {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))

	str2 := "0P"
	fmt.Println(isPalindrome(str2))
}
