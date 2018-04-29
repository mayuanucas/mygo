package problem561

import "sort"

func arrayPairSum(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i, value := range nums {
		if 0 == i%2 {
			ans += value
		}
	}
	return ans
}
