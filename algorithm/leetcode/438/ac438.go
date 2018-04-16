package problem438

func findAnagrams(s string, p string) []int {
	dict := make(map[byte]bool, len(p))
	for i := 0; i < len(p); i++ {
		dict[p[i]] = true
	}

	ans := make([]int, 0)

	for i := 0; i < len(s); i++ {
		flag := true
		for j := 0; j < len(p); j++ {
			if !dict[s[i+j]] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}

	return ans
}
