package problem594

// 和谐数组为 原数组的子数组即可，不需要是连续的子数组...
func findLHS(nums []int) int {
	dict := make(map[int]int, 0)
	for _, v := range nums {
		dict[v]++
	}

	ans := 0
	for k, v := range dict {
		if _, ok := dict[k+1]; ok {
			ans = myMax(ans, v+dict[k+1])
		}
	}
	return ans
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
