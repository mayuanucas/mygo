package problem137

func singleNumber(nums []int) int {
	dict := make(map[int]int, 0)

	for _, v := range nums {
		temp, ok := dict[v]
		if !ok {
			dict[v] = 1
		} else if 2 == temp {
			delete(dict, v)
		} else {
			dict[v]++
		}
	}
	for k := range dict {
		return k
	}

	return -1
}
