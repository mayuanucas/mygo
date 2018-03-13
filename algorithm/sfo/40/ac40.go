package problem40

import "sort"

func getLeastNumbers(data []int, k int) []int {
	if nil == data || k > len(data) {
		return nil
	}
	leastNumbers := make([]int, k)
	for i := 0; i < k; i++ {
		leastNumbers[i] = data[i]
	}
	sort.Ints(leastNumbers)

	for i:=k; i<len(data); i++ {
		if data[i] < leastNumbers[k-1] {
			leastNumbers[k-1] = data[i]
			sort.Ints(leastNumbers)
		}
	}
	return leastNumbers
}
