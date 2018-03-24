package problem069

func mySqrt(x int) int {
	if 0 == x {
		return 0
	}

	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}
