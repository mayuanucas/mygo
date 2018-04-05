package problem217

import "sort"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int, len(nums))

	for _, value := range nums {
		dict[value]++
	}

	for _, v := range dict {
		if v > 1 {
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
