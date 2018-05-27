package problem724

func pivotIndex(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return -1
	}

	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		// 左边和 == 右边和，则返回当前的下标
		if preSum[i]-nums[i] == preSum[len(nums)-1]-preSum[i] {
			return i
		}
	}
	return -1
}
