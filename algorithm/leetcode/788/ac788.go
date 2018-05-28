package problem788

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}
	return count
}

func isValid(n int) bool {
	ans := false
	for 0 < n {
		if 2 == n%10 {
			ans = true
		}
		if 5 == n%10 {
			ans = true
		}
		if 6 == n%10 {
			ans = true
		}
		if 9 == n%10 {
			ans = true
		}
		if 3 == n%10 {
			return false
		}
		if 4 == n%10 {
			return false
		}
		if 7 == n%10 {
			return false
		}
		n /= 10
	}
	return ans
}
