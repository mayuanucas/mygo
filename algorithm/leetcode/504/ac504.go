package problem504

import (
	"strconv"
)

func convertToBase7(num int) string {
	if 0 == num {
		return "0"
	}

	ans := ""
	neg := false
	if num < 0 {
		num = -num
		neg = true
	}

	for 0 != num {
		ans = strconv.Itoa(num%7) + ans
		num /= 7
	}

	if neg {
		ans = "-" + ans
	}

	return ans
}
