package problem371

func getSum(a int, b int) int {
	if 0 == b {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return getSum(sum, carry)
}
