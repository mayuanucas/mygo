package problem51

func inversePairs(data []int) int {
	if nil == data || 0 > len(data) {
		return 0
	}

	copy := make([]int, len(data))
	for id, value := range data {
		copy[id] = value
	}

	count := inversePairsCore(data, copy, 0, len(data)-1)
	return count
}

func inversePairsCore(data, copy []int, start, end int) int {
	if start == end {
		copy[start] = data[start]
		return 0
	}
	length := (end - start) / 2
	left := inversePairsCore(copy, data, start, start+length)
	right := inversePairsCore(copy, data, start+length+1, end)

	i := start + length
	j := end
	indexCopy := end
	count := 0

	for i >= start && j >= start+length+1 {
		if data[i] > data[j] {
			copy[indexCopy] = data[i]
			indexCopy--
			i--
			count += j - start - length
		} else {
			copy[indexCopy] = data[j]
			indexCopy--
			j--
		}
	}
	for ; i >= start; i-- {
		copy[indexCopy] = data[i]
		indexCopy--
	}
	for ; j >= start+length+1; j-- {
		copy[indexCopy] = data[j]
		indexCopy--
	}
	return left + right + count
}
