package problem63

func maxDiff(numbers []int) int {
	if nil == numbers || len(numbers) < 2 {
		return -1
	}

	min := numbers[0]
	max := numbers[1] - min
	for i := 2; i < len(numbers); i++ {
		if numbers[i-1] < min {
			min = numbers[i-1]
		}

		currentDiff := numbers[i] - min
		if currentDiff > max {
			max = currentDiff
		}
	}
	return max
}
