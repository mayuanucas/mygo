package problem804

func uniqueMorseRepresentations(words []string) int {
	d := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	ans := make(map[string]int, 0)
	for _, word := range words {
		ecd := ""
		for i := 0; i < len(word); i++ {
			ecd += d[word[i]-'a']
		}
		ans[ecd]++
	}
	return len(ans)
}
