package problem661

func imageSmoother(M [][]int) [][]int {
	if nil == M || len(M) <= 0 || len(M[0]) <= 0 {
		return nil
	}

	// 行数、列数
	nx, ny := len(M), len(M[0])
	ans := make([][]int, nx)
	for i := 0; i < nx; i++ {
		ans[i] = make([]int, ny)
		for j := 0; j < ny; j++ {
			ans[i][j] = smooth(M, i, j)
		}
	}
	return ans
}

func smooth(M [][]int, x, y int) int {
	nx, ny := len(M), len(M[0])
	sum, count := 0, 0

	// i: -1 -> 上一行, 0 -> 本行， 1 -> 下一行
	// j: -1 -> 左侧列, 0 -> 本列， 1 -> 右侧列
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i >= nx || y+j < 0 || y+j >= ny {
				continue
			}
			count++
			sum += M[x+i][y+j]
		}
	}
	return sum / count
}
