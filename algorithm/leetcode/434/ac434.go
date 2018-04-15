package problem434

func countSegments(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && (0 == i || ' ' == s[i-1]) {
			ans++
		}
	}
	return ans
}
