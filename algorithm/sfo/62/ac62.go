package problem62

func lastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
