package problem088

func merge(nums1 []int, m int, nums2 []int, n int) {
	idm, idn := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if idm < 0 {
			nums1[i] = nums2[idn]
			idn--
			continue
		}
		if idn < 0 {
			nums1[i] = nums1[idm]
			idm--
			continue
		}
		if nums1[idm] >= nums2[idn] {
			nums1[i] = nums1[idm]
			idm--
			continue
		}
		if nums1[idm] < nums2[idn] {
			nums1[i] = nums2[idn]
			idn--
			continue
		}
	}
}
