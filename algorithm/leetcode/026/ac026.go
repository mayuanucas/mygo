package problem026

func removeDuplicates(nums []int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	newLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[newLength-1] {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}
