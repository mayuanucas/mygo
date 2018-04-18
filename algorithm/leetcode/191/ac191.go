package problem191

func hammingWeight(n uint32) int {
	ans := 0
	for 0 != n {
		n &= n - 1
		ans++
	}
	return ans
}
