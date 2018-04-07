package problem268

import "sort"

func missingNumber(nums []int) int {
	if nil == nums {
		return -1
	}

	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == middle {
			left = middle + 1
		} else if nums[middle] > middle {
			right = middle - 1
		}
	}
	return left
}

func missingNumber2(nums []int) int {
	sum := len(nums)
	for i := 0; i < len(nums); i++ {
		sum -= i - nums[i]
	}
	return sum
}

func missingNumber3(nums []int) int {
	xor, i := 0, 0
	for ; i < len(nums); i++ {
		xor ^= i ^ nums[i]
	}
	return xor ^ i
}
