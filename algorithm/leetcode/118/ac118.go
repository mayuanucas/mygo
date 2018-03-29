package problem118

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	ans := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		oneLine := make([]int, i+1)
		oneLine[0] = 1
		oneLine[i] = 1

		for j := 1; j < i; j++ {
			oneLine[j] = ans[i-1][j-1] + ans[i-1][j]
		}

		ans = append(ans, oneLine)
	}
	return ans
}
