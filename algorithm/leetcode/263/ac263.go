package problem263

func isUgly(num int) bool {
	for i := 5; i >= 2 && num > 0; i-- {
		for 0 == num%i {
			num /= i
		}
	}
	return 1 == num
}
