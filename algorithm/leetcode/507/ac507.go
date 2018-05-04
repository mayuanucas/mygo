package problem507

import "math"

func checkPerfectNumber(num int) bool {
	if 1 == num {
		return false
	}

	// 1 为该数字的正因子，需要计入
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if 0 == num%i {
			sum += i + num/i
		}
	}

	return sum == num
}
