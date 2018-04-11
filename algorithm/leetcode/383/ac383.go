package problem383

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 && len(magazine) == 0 {
		return true
	}
	if len(magazine) == 0 {
		return false
	}
	if len(ransomNote) == 0 {
		return true
	}

	dict1 := make(map[rune]int, 0)
	for _, v := range ransomNote {
		dict1[v]++
	}

	dict2 := make(map[rune]int, 0)
	for _, v := range magazine {
		dict2[v]++
	}

	for k1, v1 := range dict1 {
		if v2, ok := dict2[k1]; !ok || v2 < v1 {
			return false
		}
	}
	return true
}
