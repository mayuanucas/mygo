package problem058

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	subArray := strings.Split(s, " ")
	for i := len(subArray) - 1; i >= 0; i-- {
		length := len(subArray[i])
		if length > 0 {
			return length
		}
	}
	return 0
}

func lengthOfLastWord2(s string) int {
	len, tail := 0, len(s)-1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}
	for tail >= 0 && s[tail] != ' ' {
		len++
		tail--
	}
	return len
}
