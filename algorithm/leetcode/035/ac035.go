package problem035

func searchInsert(nums []int, target int) int {
	if nil == nums || len(nums) < 1 {
		return 0
	}

	start, end := 0, len(nums)-1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return start
}
