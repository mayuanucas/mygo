package problem557

func reverseWords(s string) string {
	strChar := []rune(s)

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			end = i - 1
			reverse(strChar, start, end)
			start = i + 1
		}
	}
	end = len(s)-1
	reverse(strChar, start, end)
	return string(strChar)
}

func reverse(chars []rune, start, end int) {
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}
