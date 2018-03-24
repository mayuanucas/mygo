package problem067

func addBinary(a string, b string) string {
	aChar := []rune(a)
	bChar := []rune(b)
	var s []rune
	var c rune
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || c == 1 {
		if i >= 0 {
			c += aChar[i] - '0'
			i--
		}
		if j >= 0 {
			c += bChar[j] - '0'
			j--
		}
		s = append(s, c%2+'0')
		c /= 2
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}

func addBinary2(a string, b string) string {
	s := make([]byte, 0)
	var c byte
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || 1 == c {
		if i >= 0 {
			c += a[i] - '0'
			i--
		}
		if j >= 0 {
			c += b[j] - '0'
			j--
		}
		s = append(s, c%2+'0')
		c /= 2
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
