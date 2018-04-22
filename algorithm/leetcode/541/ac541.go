package problem541

func reverseStr(s string, k int) string {
	chars := []byte(s)

	for idx := 0; idx*k < len(chars); idx += 2 {
		start, end := idx*k, idx*k+k-1
		if end >= len(chars) {
			end = len(chars)-1
		}

		for start < end {
			chars[start], chars[end] = chars[end], chars[start]
			start++
			end--
		}
	}
	return string(chars)
}
