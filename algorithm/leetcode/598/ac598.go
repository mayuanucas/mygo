package problem598

func maxCount(m int, n int, ops [][]int) int {
	if nil == ops || len(ops) == 0 {
		return m * n
	}

	row, col := ops[0][0], ops[0][1]
	for _, v := range ops {
		row = myMin(row, v[0])
		col = myMin(col, v[1])
	}
	return row * col
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
