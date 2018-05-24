package problem693

import (
	"strings"
	"fmt"
)

func hasAlternatingBits(n int) bool {
	res := decimalToBinary(n)
	for i, temp := 1, res[0]; i < len(res); i++ {
		if temp == res[i] {
			return false
		} else {
			temp = res[i]
		}
	}
	return true
}

func decimalToBinary(n int) []int {
	if 0 == n {
		return []int{0}
	}

	ans := make([]int, 0)
	for 0 != n {
		ans = append(ans, n%2)
		n /= 2
	}
	return ans
}

func hasAlternatingBits2(n int) bool {
	strN := fmt.Sprintf("%b", n)
	return !(strings.Contains(strN, "00") || strings.Contains(strN, "11"))
}
