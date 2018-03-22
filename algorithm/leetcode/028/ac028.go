package problem028

func strStr(haystack string, needle string) int {
	if "" == needle {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle) && haystack[i+j] == needle[j]; j++ {
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}
