package problem189

func rotate(nums []int, k int) {
	if nil == nums {
		return
	}

	length := len(nums)
	if k > length {
		k %= length
	}
	if 0 == k || k == length {
		return
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
