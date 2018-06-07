package problem840

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 { return 0 }
	res := 0
	for i:=1; i<len(grid)-1; i++ {
		for j:=1; j<len(grid[0])-1; j++ {
			if grid[i][j] == 5 {
				is_ok := true
				m := make([]bool, 9)
				m[4] = true

				for _, offset := range [][]int{[]int{-1, -1}, []int{-1, 0},
					[]int{-1, 1}, []int{0, -1}} {
					v1,v2 := grid[i+offset[0]][j+offset[1]], grid[i-offset[0]][j-offset[1]]
					if v1<1 || v1>9 || v2<1 || v2>9 ||
						v1+v2 != 10 || m[v1-1] == true || m[v2-1] == true {
						is_ok = false
						break
					}
					m[v1-1], m[v2-1] = true, true
				}
				if is_ok == true { res++ }
			}
		}
	}
	return res
}
