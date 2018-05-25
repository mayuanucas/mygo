package problem695

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if 1 == grid[i][j] {
				current := AreaOfIsland(grid, i, j)
				if maxArea < current {
					maxArea = current
				}
			}
		}
	}
	return maxArea
}

func AreaOfIsland(grid [][]int, i, j int) int {
	if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && 1 == grid[i][j] {
		// 当前元素置为 0 ，避免重复计算
		grid[i][j] = 0
		return 1 + AreaOfIsland(grid, i-1, j) + AreaOfIsland(grid, i+1, j) + AreaOfIsland(grid, i, j-1) + AreaOfIsland(grid, i, j+1)
	}
	return 0
}
