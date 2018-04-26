package problem520

func detectCapitalUse(word string) bool {
	if len(word) < 1 {
		return false
	} else if len(word) == 1 {
		return true
	}

	firstChar := word[0]
	secondChar := word[1]
	if 'a' <= firstChar && firstChar <= 'z' {
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
	} else if 'a' <= secondChar && secondChar <= 'z' {
		for i := 2; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] < 'A' || word[i] > 'Z' {
				return false
			}
		}
	}
	return true
}
