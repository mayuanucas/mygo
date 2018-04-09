package problem303

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{data: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if 0 == i {
		return this.data[j]
	}
	return this.data[j] - this.data[i-1]
}
