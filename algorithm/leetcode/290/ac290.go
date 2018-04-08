package problem290

import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	dictPattern := make(map[byte]int)
	dictString := make(map[string]int)
	for i := 0; i < len(pattern); i++ {
		patternValue := dictPattern[pattern[i]]
		dictPattern[pattern[i]] = i+1

		stringValue := dictString[words[i]]
		dictString[words[i]] = i+1

		if patternValue != stringValue {
			return false
		}
	}
	return true
}
