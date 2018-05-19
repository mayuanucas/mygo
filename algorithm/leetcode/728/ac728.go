package problem728

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for i := left; i <= right; i++ {
		value := i
		for ; value > 0; value /= 10 {
			if 0 == value%10 || 0 != (i%(value%10)) {
				break
			}
		}
		if 0 == value {
			ans = append(ans, i)
		}
	}
	return ans
}
