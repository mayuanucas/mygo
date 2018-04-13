package problem414

import (
	"sort"
)

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

func thirdMax2(nums []int) int {
	empty := [4]bool{true, true, true, true}
	max := [4]int{}
	for _, v := range nums {
		if (!empty[1] && v == max[1]) || (!empty[2] && v == max[2]) || (!empty[3] && v == max[3]) {
			continue
		}
		if empty[1] || v > max[1] {
			max[3] = max[2]
			max[2] = max[1]
			max[1] = v

			if !empty[2] {
				empty[3] = false
			}
			if !empty[1] {
				empty[2] = false
			}
			empty[1] = false
		} else if empty[2] || v > max[2] {
			max[3] = max[2]
			max[2] = v

			if !empty[2] {
				empty[3] = false
			}
			empty[2] = false
		} else if empty[3] || v > max[3] {
			max[3] = v

			empty[3] = false
		}
	}

	if !empty[3] {
		return max[3]
	} else {
		return max[1]
	}
}
