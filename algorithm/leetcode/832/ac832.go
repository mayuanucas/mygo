package problem832

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		reverse(A[i])
		invert(A[i])
	}
	return A
}

func reverse(array []int) {
	for left, right := 0, len(array)-1; left < right; {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}

func invert(array []int) {
	for i := 0; i < len(array); i++ {
		if 1 == array[i] {
			array[i] = 0
		} else if 0 == array[i] {
			array[i] = 1
		}
	}
}
