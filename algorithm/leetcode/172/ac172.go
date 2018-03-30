package problem172

func trailingZeroes(n int) int {
	if 0 == n {
		return 0
	} else {
		return n/5 + trailingZeroes(n/5)
	}
}
