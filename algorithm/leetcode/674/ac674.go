package problem674

func findLengthOfLCIS(nums []int) int {
	if nil == nums {
		return 0
	}

	ans, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		if 0 == i || nums[i-1] < nums[i] {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 1
		}
	}
	return ans
}
