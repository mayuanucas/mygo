package problem292

func canWinNim(n int) bool {
	if n < 1 {
		return false
	}

	return 0 != n % 4
}
