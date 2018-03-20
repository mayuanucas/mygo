package problem65

func Add(num1, num2 int) int {
	var sum, carry int
	for {
		sum = num1 ^ num2
		carry = (num1 & num2) << 1
		num1 = sum
		num2 = carry

		if 0 == num2 {
			break
		}
	}
	return num1
}
