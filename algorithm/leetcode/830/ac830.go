package problem830

func largeGroupPositions(S string) [][]int {
	if len(S) < 3 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	start, count := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == S[start] {
			count++
		} else {
			if 3 <= count {
				ans = append(ans, []int{start, i - 1})
			}
			start = i
			count = 1
		}
	}

	// large group 出现在字符串末尾位置
	if 3 <= count && S[start] == S[len(S)-1] {
		ans = append(ans, []int{start, len(S) - 1})
	}

	return ans
}
