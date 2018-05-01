package problem566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nil == nums || 0 == len(nums) || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	temp := make([]int, 0)
	for row := 0; row < len(nums); row++ {
		for col := 0; col < len(nums[0]); col++ {
			temp = append(temp, nums[row][col])
		}
	}

	ans := make([][]int, r)
	for row := 0; row < r; row++ {
		tmp := make([]int, c)
		for col := 0; col < c; col++ {
			tmp[col] = temp[row*c+col]
		}
		ans[row] = tmp
	}
	return ans
}

func matrixReshape2(nums [][]int, r int, c int) [][]int {
	if nil == nums || 0 == len(nums) || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	originCol := len(nums[0])

	ans := make([][]int, r)
	for row := 0; row < r; row++ {
		temp := make([]int, c)
		for col := 0; col < c; col++ {
			idx := row*c + col
			currentRow := idx / originCol
			currentCol := idx % originCol
			temp[col] = nums[currentRow][currentCol]
		}
		ans[row] = temp
	}
	return ans
}
