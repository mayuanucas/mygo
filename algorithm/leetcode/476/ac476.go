package problem476

func findComplement(num int) int {
	temp, mask := num, 0

	for temp > 0 {
		mask = mask*2 + 1
		temp /= 2
	}
	return mask ^ num
}
