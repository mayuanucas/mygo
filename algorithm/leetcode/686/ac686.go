package problem686

import "strings"

func repeatedStringMatch(A string, B string) int {
	if strings.Contains(A, B) {
		return 1
	}

	tempA := A
	for i := 0; len(tempA) <= 10000; i++ {
		tempA += A
		if strings.Contains(tempA, B) {
			return i + 2
		}
	}
	return -1
}
