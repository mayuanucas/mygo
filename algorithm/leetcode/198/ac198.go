package problem198

func rob(nums []int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	if 1 == len(nums) {
		return nums[0]
	} else if 2 == len(nums) {
		return maxInt(nums[0], nums[1])
	}

	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = maxInt(nums[i]+f[i-2], f[i-1])
	}

	return f[len(nums)-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
