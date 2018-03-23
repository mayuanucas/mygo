package problem038

import "strconv"

func countAndSay(n int) string {
	if 0 == n {
		return ""
	}

	ans := "1"
	for n--; n > 0; n-- {
		current := ""
		for i := 0; i < len(ans); i++ {
			count := 1
			for i+1 < len(ans) && ans[i] == ans[i+1] {
				count++
				i++
			}
			current += strconv.Itoa(count) + string(ans[i])
		}
		ans = current
	}
	return ans
}
