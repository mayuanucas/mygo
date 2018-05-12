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
