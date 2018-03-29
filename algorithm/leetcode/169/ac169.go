package problem169

func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	count, flag := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if count < 1 {
			flag = nums[i]
			count = 1
			continue
		}
		if nums[i] == flag {
			count++
		} else {
			count--
		}
	}
	return flag
}
