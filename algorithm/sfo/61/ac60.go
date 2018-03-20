package problem61

import "sort"

func isContinuous(numbers []int) bool {
	if nil == numbers || len(numbers) < 5 {
		return false
	}

	sort.Ints(numbers)
	numberOfZero, numberOfGap := 0, 0
	for _, value := range numbers {
		if value == 0 {
			numberOfZero++
		}
	}

	small, big := 0, 1
	for big < len(numbers) {
		if numbers[small] == numbers[big] {
			return false
		}
		numberOfGap += numbers[big] - numbers[small] - 1
		small = big
		big++
	}
	if numberOfGap > numberOfZero {
		return false
	}
	return true
}
