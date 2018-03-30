package problem171

func titleToNumber(s string) int {
	ans := 0
	for _, value := range s {
		ans = int(value) - int('A') + 1 + ans*26
	}
	return ans
}
