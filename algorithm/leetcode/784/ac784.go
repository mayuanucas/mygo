package problem784

// 递归解法
// 每当遇到一个字符时，分别统计 1.该字符为小写时的所有情况 2.该字符为大写时的所有情况
func letterCasePermutation(S string) []string {
	ans := make([]string, 0)
	if len(S) < 1 {
		return []string{""}
	}
	return core(ans, []byte(S), 0)
}

func core(ans []string, chars []byte, index int) []string {
	if index == len(chars) {
		ans = append(ans, string(chars))
	} else {
		if isLetter(chars[index]) {
			chars[index] = toLower(chars[index])
			ans = core(ans, chars, index+1)
			chars[index] = toUpper(chars[index])
		}
		ans = core(ans, chars, index+1)
	}
	return ans
}

func isLetter(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		dis := byte('a' - 'A')
		return c + dis
	}
	return c
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		dis := byte('a' - 'A')
		return c - dis
	}
	return c
}
