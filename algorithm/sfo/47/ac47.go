package problem47

func getMaxValue(matrix [][]int) int {
	if nil == matrix {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			if 0 == i && 0 == j {
				dp[i][j] = matrix[i][j]
			}
			if 0 == i && 0 != j {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			}
			if 0 != i && 0 == j {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			}
			if 0 != i && 0 != j {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + matrix[i][j]
				}
			}
		}
	}
	return dp[rows-1][columns-1]
}

func getMaxValue2(matrix [][]int) int {
	if nil == matrix {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	dp := make([]int, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			left, up := 0, 0
			if 0 < i {
				up = dp[j]
			}
			if 0 < j {
				left = dp[j-1]
			}

			if left >= up {
				dp[j] = left + matrix[i][j]
			} else {
				dp[j] = up + matrix[i][j]
			}
		}
	}
	return dp[columns-1]
}
