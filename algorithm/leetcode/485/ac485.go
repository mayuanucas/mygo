package problem485

func findMaxConsecutiveOnes(nums []int) int {
	ans, count := 0, 0
	for _, v := range nums {
		if 1 == v {
			count++
			if ans < count {
				ans = count
			}
		} else {
			count = 0
		}
	}
	return ans
}
