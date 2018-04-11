package problem349

func intersection(nums1 []int, nums2 []int) []int {
	if nil == nums1 || len(nums1) == 0 {
		return nums1
	}
	if nil == nums2 || len(nums2) == 0 {
		return nums2
	}

	dict1 := make(map[int]int, 1)
	dict2 := make(map[int]int, 1)
	for _, v := range nums1 {
		dict1[v]++
	}
	for _, v := range nums2 {
		dict2[v]++
	}

	ans := make([]int, 0)
	for k := range dict1 {
		if _, ok := dict2[k]; ok {
			ans = append(ans, k)
		}
	}
	return ans
}

func intersection2(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	for _, v := range nums1 {
		dict[v] = 1
	}

	for _, v := range nums2 {
		if 1 == dict[v] {
			dict[v] = 2
		}
	}

	ans := make([]int, 0)
	for k, v := range dict {
		if 2 == v {
			ans = append(ans, k)
		}
	}
	return ans
}
