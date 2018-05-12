package problem532

func findPairs(nums []int, k int) int {
	if nil == nums || len(nums) == 0 || k < 0 {
		return 0
	}

	dict := make(map[int]int, 0)
	for _, v := range nums {
		dict[v]++
	}

	ans := 0
	for key, value := range dict {
		if 0 == k {
			if value >= 2 {
				ans++
			}
		} else {
			if _, ok := dict[key+k]; ok {
				ans++
			}
		}
	}
	return ans
}
