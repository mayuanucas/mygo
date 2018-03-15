package problem46

import "strconv"

func getTranslation(number int) int {
	if 0 > number {
		return 0
	}
	return getTranslationCore(strconv.Itoa(number))
}

func getTranslationCore(str string) int {
	length := len(str)
	counts := make([]int, length)
	num := 0

	for i := length - 1; i >= 0; i-- {
		num = 0
		if i < length-1 {
			num = counts[i+1]
		} else {
			num = 1
		}

		if i < length-1 {
			digit1 := str[i] - '0'
			digit2 := str[i+1] - '0'
			number := digit1*10 + digit2
			if 10 <= number && 25 >= number {
				if i < length-2 {
					num += counts[i+2]
				} else {
					num += 1
				}
			}
		}
		counts[i] = num
	}
	return counts[0]
}
