package problem125

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isCharacter(s[left]) {
			left++
		}
		for left < right && !isCharacter(s[right]) {
			right--
		}

		if lowerCharacter(s[left]) != lowerCharacter(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isCharacter(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	} else {
		return false
	}
}

func lowerCharacter(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}
