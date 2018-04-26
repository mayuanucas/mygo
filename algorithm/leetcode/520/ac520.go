package problem520

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
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

// 直接检测单词中大写字母的数量
// 符合规范的有下面几种情况：
// 1. 全部都是大写单词 -> 大写单词数目 = 单词长度
// 2. 没有大写单词 -> 大写单词数目 = 0
// 3. 只有头部是大写单词 -> 大写单词数目 = 1 && 仅头部第一个单词是大写
func detectCapitalUse2(word string) bool {
	if len(word) <= 1 {
		return true
	}

	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] < 'a' {
			count++
		}
	}

	if 0 == count || len(word) == count || (1 == count && word[0] < 'a') {
		return true
	}
	return false
}
