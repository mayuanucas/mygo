package problem387

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	dict := make(map[rune]int, 0)
	for _, v := range s {
		dict[v]++
	}

	ans := -1
	for id, v := range s {
		if 1 == dict[v] {
			return id
		}
	}
	return ans
}

// 效率更高
func firstUniqChar2(s string) int {
	position := [26]int{}

	for _, v := range s {
		position[v-'a']++
	}
	for id, v := range s {
		if 1 == position[v-'a'] {
			return id
		}
	}
	return -1
}
