package problem448

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		id := abs(nums[i]) - 1
		if nums[id] > 0 {
			nums[id] = -nums[id]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
