package problem706

import "strings"

// 拼接2个字符串，正好满足了循环条件
func rotateString(A string, B string) bool {
	if len(A) != len(B){
		return false
	}
	return strings.Contains(A+A, B)
}
