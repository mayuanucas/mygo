package problem844

func backspaceCompare(S string, T string) bool {
	newS := core(S)
	newT := core(T)

	if newS == newT {
		return true
	}
	return false
}

func core(s string) string {
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if '#' == s[i] {
			if 0 < len(ans) {
				ans = ans[0 : len(ans)-1]
			}
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
