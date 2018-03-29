package problem119

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	vi := make([]int, rowIndex+1)
	vi[0] = 1
	for i := 0; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			vi[j] = vi[j] + vi[j-1]
		}
	}
	return vi
}
