package problem819

import (
	"strings"
	"regexp"
)

func mostCommonWord(paragraph string, banned []string) string {
	text := strings.ToLower(paragraph)

	reg := regexp.MustCompile("[a-z]+")
	result := reg.FindAllString(text, -1)

	ans := make(map[string]int, 0)

	for _, word := range result {
		if !isContain(banned, word) {
			ans[word]++
		}
	}

	maxTime := 0
	var ansString string
	for k, v := range ans {
		if v > maxTime {
			maxTime = v
			ansString = k
		}
	}
	return ansString
}

func isContain(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
