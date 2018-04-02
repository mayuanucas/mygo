package problem202

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)

		if 1 == fast {
			return true
		}
		if slow == fast {
			break
		}
	}
	return false
}

func digitSquareSum(n int) int {
	sum, tmp := 0, 0
	for ; 0 != n; n /= 10 {
		tmp = n % 10
		sum += tmp * tmp
	}
	return sum
}
