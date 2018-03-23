package problem053

func maxSubArray(nums []int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	preSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if preSum+nums[i] <= nums[i] {
			preSum = nums[i]
		} else {
			preSum += nums[i]
		}
		if preSum > maxSum {
			maxSum = preSum
		}
	}
	return maxSum
}
