package problem665

func checkPossibility(nums []int) bool {
	if nil == nums || len(nums) <= 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	count := 0
	for i := 1; i < len(nums) && count <= 1; i++ {
		if nums[i-1] > nums[i] {
			count++
			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}

	return count <= 1
}
