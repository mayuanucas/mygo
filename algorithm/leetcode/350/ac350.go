package problem350

func intersect(nums1 []int, nums2 []int) []int {
	if nil == nums1 || len(nums1) == 0 {
		return nums1
	}
	if nil == nums2 || len(nums2) == 0 {
		return nums2
	}

	dict1 := make(map[int]int, 0)
	for _, v := range nums1 {
		dict1[v]++
	}

	dict2 := make(map[int]int, 0)
	for _, v := range nums2 {
		dict2[v]++
	}

	ans := make([]int, 0)
	for k1, v1 := range dict1 {
		if v2, ok := dict2[k1]; ok {
			for repeat := minInt(v1, v2); repeat > 0; repeat-- {
				ans = append(ans, k1)
			}
		}
	}
	return ans
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
