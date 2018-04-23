package problem120

// 自下向上推导
// ans数组保存的是：到达下一行某位置的最小路径和
func minimumTotal(triangle [][]int) int {
	ans := make([]int, len(triangle)+1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			ans[j] = myMin(ans[j], ans[j+1]) + triangle[i][j]
		}
	}

	return ans[0]
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
