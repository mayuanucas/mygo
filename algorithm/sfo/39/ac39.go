package problem39

func moreThanHalfNum(numbers []int) int {
	if nil == numbers || 0 >= len(numbers) {
		return 0
	}
	ans := numbers[0]
	times := 1
	for i := 1; i < len(numbers); i++ {
		if 0 == times {
			ans = numbers[i]
			times = 1
		} else if ans == numbers[i] {
			times++
		} else {
			times--
		}
	}
	return ans
}
