package problem581

import "sort"

func findUnsortedSubarray(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	numsBackup := make([]int, len(nums))
	copy(numsBackup, nums)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for ; left < len(nums) && nums[left] == numsBackup[left]; {
		left++
	}
	// 该数组是有序数组
	if left == len(nums) {
		return 0
	}

	for ; right >= 0 && nums[right] == numsBackup[right]; {
		right--
	}
	return right - left + 1
}

func findUnsortedSubarray2(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	begin, end := -1, -2
	min, max := nums[n-1], nums[0]
	for i := 1; i < n; i++ {
		max = myMax(max, nums[i])
		min = myMin(min, nums[n-1-i])
		if nums[i] < max {
			end = i
		}
		if nums[n-1-i] > min {
			begin = n - 1 - i
		}
	}
	return end - begin + 1
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
