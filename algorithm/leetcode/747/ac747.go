package problem747

func dominantIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	id1, id2 := -1, -1
	for i := 0; i < len(nums); i++ {
		if -1 == id1 || nums[i] > nums[id1] {
			id2 = id1
			id1 = i
		} else if -1 == id2 || nums[i] > nums[id2] {
			id2 = i
		}
	}

	if nums[id2]*2 <= nums[id1] {
		return id1
	}
	return -1
}
