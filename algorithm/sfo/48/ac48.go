package problem48

func longestSubstring(str string) int {
	curLength, maxLength := 0, 0
	position := make([]int, 26)
	for i := 0; i < 26; i++ {
		position[i] = -1
	}

	for i := 0; i < len(str); i++ {
		preIndex := position[str[i]-'a']
		if 0 > preIndex || i-preIndex > curLength {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = i - preIndex
		}
		position[str[i]-'a'] = i
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}
