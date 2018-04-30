package problem561

import "sort"

func arrayPairSum(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i, value := range nums {
		if 0 == i%2 {
			ans += value
		}
	}
	return ans
}

func arrayPairSum2(nums []int) int {
	//给数组排序，然后挨着的两个为一组，也就是取第1,3,....到2n-3,2n-1的和(从1开始计数)
	dict := make([]int, 20001)
	//考虑负数
	for _, v := range nums {
		dict[v+10000]++
	}

	id, count := 0, 0
	for k, v := range dict {
		//考虑有相同数据
		for v > 0 {
			id++
			if 1 == id%2 {
				count += k - 10000
			}
			v--
		}
	}
	return count
}
