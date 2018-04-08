package problem283

func moveZeroes(nums []int) {
	if nil == nums || len(nums) == 0 {
		return
	}

	insertPos := 0
	for _, values := range nums {
		if 0 != values {
			nums[insertPos] = values
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
