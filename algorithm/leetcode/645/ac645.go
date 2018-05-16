package problem645

func findErrorNums(nums []int) []int {
	if nil == nums || len(nums) < 2 {
		return nil
	}

	ans := make([]int, 0)
	for _, v := range nums {
		if 0 > nums[myAbs(v)-1] {
			// 找到重复的数字
			ans = append(ans, myAbs(v))
		} else {
			nums[myAbs(v)-1] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		// 找到缺失元素的位置
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
