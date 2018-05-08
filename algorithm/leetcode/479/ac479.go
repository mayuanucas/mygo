package problem479

import (
	"strconv"
	"math"
)

func largestPalindrome(n int) int {
	if 1 == n {
		return 9
	}

	max := int(math.Pow10(n) - 1)
	for v := max - 1; v > max/10; v-- {
		u, _ := strconv.Atoi(strconv.Itoa(v) + reverseInt(v))
		for x := max; x*x >= u; x-- {
			if u%x == 0 {
				return u % 1337
			}
		}
	}
	return 0
}

func reverseInt(a int) string {
	temp := []byte(strconv.Itoa(a))
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return string(temp)
}
