package problem344

func reverseString(s string) string {
	char := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		char[i], char[j] = char[j], char[i]
	}
	return string(char)
}
