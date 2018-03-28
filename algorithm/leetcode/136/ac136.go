package problem136

func singleNumber(nums []int) int {
	ans := 0
	for _, value := range nums{
		ans ^= value
	}
	return ans
}
