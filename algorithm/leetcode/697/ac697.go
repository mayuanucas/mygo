package problem697

func findShortestSubArray(nums []int) int {
	size := len(nums)
	if 2 > size {
		return size
	}

	first := make(map[int]int, size)
	count := make(map[int]int, size)
	maxCount := 1
	minLen := size
	for i, v := range nums {
		count[v]++
		if 1 == count[v] {
			first[v] = i
		} else {
			l := i - first[v] + 1
			if maxCount < count[v] || (maxCount == count[v] && l < minLen) {
				maxCount = count[v]
				minLen = l
			}
		}
	}

	if len(count) == size {
		return 1
	}
	return minLen
}
