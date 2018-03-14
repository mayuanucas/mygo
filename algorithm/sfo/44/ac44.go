package problem44

import "math"

func countOfIntegers(digits int) int {
	if 1 == digits {
		return 10
	}
	return int(9 * math.Pow10(digits-1))
}

func beginNumber(digits int) int {
	if 1 == digits {
		return 0
	}
	return int(math.Pow10(digits - 1))
}

func digitAtIndexCore(index, digits int) int {
	number := beginNumber(digits) + index/digits
	indexFromRight := digits - index%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func digitAtIndex(index int) int {
	if 0 > index {
		return -1
	}
	digit := 1
	for {
		numbers := countOfIntegers(digit)
		if index < numbers*digit {
			return digitAtIndexCore(index, digit)
		}
		index -= digit * numbers
		digit++
	}
	return -1
}
