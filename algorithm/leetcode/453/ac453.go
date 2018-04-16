package problem453

func minMoves(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	minNumber := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < minNumber {
			minNumber = nums[i]
		}
	}

	return sum - minNumber*len(nums)
}
