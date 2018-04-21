package problem496

func nextGreaterElement(findNums []int, nums []int) []int {
	ans := make([]int, 0)
	for _, v := range findNums {
		ans = append(ans, nextMax(v, nums))
	}
	return ans
}

func nextMax(n int, nums []int) int {
	ans := -1
	index := 0

	// 找到数字 n 在 nums 切片中的位置
	for index < len(nums) && nums[index] != n {
		index++
	}

	for index < len(nums) {
		if nums[index] > n {
			ans = nums[index]
			break
		}
		index++
	}
	return ans
}
