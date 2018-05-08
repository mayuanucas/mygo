package problem475

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	i, ans := 0, 0
	for _, house := range houses {
		for i < len(heaters)-1 && heaters[i]+heaters[i+1] <= house*2 {
			i++
		}
		ans = maxInt(ans, absInt(heaters[i]-house))
	}
	return ans
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
