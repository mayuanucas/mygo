package problem461

func hammingDistance(x int, y int) int {
	n := x ^ y
	ans := 0
	for n != 0 {
		n &= n - 1
		ans++
	}
	return ans
}
