package problem463

func islandPerimeter(grid [][]int) int {
	isLand, neighbours := 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if 1 == grid[i][j] {
				isLand++
				// 统计下一行与该陆地块相邻的数量
				if i < len(grid)-1 && 1 == grid[i+1][j] {
					neighbours++
				}
				// 统计右侧与该陆地块相邻的数量
				if j < len(grid[i]) - 1 && 1 == grid[i][j+1] {
					neighbours++
				}
			}
		}
	}
	// and the pattern is islands * 4 - neighbours * 2, since every adjacent islands made two sides disappeared
	// 4 + 4 - ? = 6  -> ? = 2
	return isLand*4 - neighbours*2
}
