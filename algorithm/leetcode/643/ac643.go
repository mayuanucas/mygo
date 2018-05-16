package problem643

import "math"

func findMaxAverage(nums []int, k int) float64 {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	maxSum := math.MinInt32
	for i := 0; i <= len(nums)-k; i++ {
		temp := 0
		for j := 0; j < k; j++ {
			temp += nums[i+j]
		}
		if maxSum < temp {
			maxSum = temp
		}
	}
	return float64(maxSum) / float64(k)
}
