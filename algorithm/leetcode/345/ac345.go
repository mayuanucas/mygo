package problem345

import "strings"

func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	vowels := "aeiouAEIOU"
	chars := []rune(s)

	for start, end := 0, len(chars)-1; start < end; {
		for start < end && !strings.ContainsRune(vowels, chars[start]) {
			start++
		}

		for start < end && !strings.ContainsRune(vowels, chars[end]) {
			end--
		}

		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
	return string(chars)
}
