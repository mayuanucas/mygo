package problem168

func convertToTitle(n int) string {
	ans := make([]byte, 0)

	for n > 0 {
		n--
		ans = append(ans, byte(n%26)+'A')
		n /= 26
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
