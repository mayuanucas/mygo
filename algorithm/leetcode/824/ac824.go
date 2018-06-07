package problem824

import (
	"strings"
	"bytes"
)

func toGoatLatin(S string) string {
	words := strings.Fields(S)
	var buffer bytes.Buffer

	vowels := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
		"A": 1,
		"E": 1,
		"I": 1,
		"O": 1,
		"U": 1,
	}

	for i := 0; i < len(words); i++ {
		first := string(words[i][0])
		if vowels[first] == 1 {
			buffer.WriteString(words[i])
			buffer.WriteString("ma")
		} else {
			buffer.WriteString(string(words[i][1:]))
			buffer.WriteString(first)
			buffer.WriteString("ma")
		}
		for j := 0; j <= i; j++ {
			buffer.WriteString("a")
		}
		if i != len(words)-1 {
			buffer.WriteString(" ")
		}

	}
	return buffer.String()
}
