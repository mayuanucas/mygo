package problem628

import (
	"sort"
	"math"
)

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	ans1, ans2 := 1, 1
	for i := len(nums) - 1; i >= len(nums)-3; i-- {
		ans1 *= nums[i]
	}

	for i := 0; i < 2; i++ {
		ans2 *= nums[i]
	}
	ans2 *= nums[len(nums)-1]

	if ans1 > ans2 {
		return ans1
	}
	return ans2
}

func maximumProduct2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}
	return myMax(max1*max2*max3, min1*min2*max1)
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
