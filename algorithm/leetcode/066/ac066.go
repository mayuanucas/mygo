package problem066

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if 9 != digits[i] {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if 0 == digits[0] {
		newAns := make([]int, len(digits)+1)
		newAns[0] = 1
		return newAns
	}
	return digits
}
