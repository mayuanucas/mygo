package problem680

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return isPalindromic(s, l+1, r) || isPalindromic(s, l, r-1)
		} else {
			l++
			r--
		}
	}
	return true
}

func isPalindromic(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
