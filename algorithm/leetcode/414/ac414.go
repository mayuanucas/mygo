package problem414

import "sort"

func thirdMax(nums []int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}

	sort.Ints(nums)

	ans := nums[length-1]
	j := 2
	for i := length - 2; i >= 0 && j > 0; i-- {
		if nums[i] != ans {
			ans = nums[i]
			j--
		}
	}

	if 0 != j {
		ans = nums[length-1]
	}
	return ans
}
