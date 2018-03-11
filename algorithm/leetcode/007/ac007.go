package problem007

import "math"

func reverse(x int) int {
	var result int64
	for 0 != x {
		result = result*10 + int64(x%10)
		x /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return int(result)
}
