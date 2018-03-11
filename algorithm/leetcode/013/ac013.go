package problem013

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) <= 0 {
		return 0
	}
	// 当 map 中不存在该键时 ok 为 false
	sum, ok := roman[s[len(s)-1]]
	if !ok {
		return 0
	}

	for i := len(s) - 2; i >= 0; i-- {
		if roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}
	return sum
}
