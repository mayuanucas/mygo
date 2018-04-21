package problem500

import "strings"

func findWords(words []string) []string {
	ans := make([]string, 0)
	data := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

	for _, v := range words {
		word := strings.ToLower(v)
		if (strings.ContainsAny(data[0], word) && !strings.ContainsAny(data[1], word) && !strings.ContainsAny(data[2], word)) ||
			(!strings.ContainsAny(data[0], word) && strings.ContainsAny(data[1], word) && !strings.ContainsAny(data[2], word)) ||
			(!strings.ContainsAny(data[0], word) && !strings.ContainsAny(data[1], word) && strings.ContainsAny(data[2], word)) {
			ans = append(ans, v)
		}
	}
	return ans
}
