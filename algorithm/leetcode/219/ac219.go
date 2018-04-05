package problem219

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if value, ok := dict[nums[i]]; ok {
			if -k <= value-i || value-i <= k {
				return true
			}
		}
		dict[nums[i]] = i
	}

	return false
}
