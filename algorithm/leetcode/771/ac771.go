package problem771

func numJewelsInStones(J string, S string) int {
	ans := 0
	dict := make(map[uint8]bool, len(J))
	for i := 0; i < len(J); i++ {
		dict[J[i]] = true
	}
	for i := 0; i < len(S); i++ {
		if dict[S[i]] {
			ans++
		}
	}
	return ans
}
