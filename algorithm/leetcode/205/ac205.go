package problem205

func isIsomorphic(s string, t string) bool {
	length := len(s)
	num1 := make([]int, 256)
	num2 := make([]int, 256)

	for i := 0; i < length; i++ {
		if num1[s[i]] != num2[t[i]] {
			return false
		}
		num1[s[i]] = i + 1
		num2[t[i]] = i + 1
	}
	return true
}
