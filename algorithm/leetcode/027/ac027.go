package problem027

func removeElement(nums []int, val int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	var idx int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
