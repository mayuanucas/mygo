package problem014

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if nil == strs || 0 >= len(strs) {
		return ""
	}
	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]
	var i, length int
	if len(first) < len(last) {
		length = len(first)
	} else {
		length = len(last)
	}

	for i < length && first[i] == last[i] {
		i++
	}
	return first[0:i]
}
