package problem367

func isPerfectSquare(num int) bool {
	if 1 == num {
		return true
	}

	for middle := num / 2; middle >= 2; middle-- {
		result := middle * middle

		if result == num {
			return true
		} else if result < num {
			return false
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	low, high := 1, num

	for low <= high {
		middle := (low + high) >> 1
		result := middle * middle

		if result == num {
			return true
		} else if result < num {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false
}
