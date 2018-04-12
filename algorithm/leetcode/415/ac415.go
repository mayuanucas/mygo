package problem415

func addStrings(num1 string, num2 string) string {
	str1 := []rune(num1)
	str2 := []rune(num2)
	ans := make([]rune, 0)
	var carry rune = 0
	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || 1 == carry; i, j = i-1, j-1 {
		var x, y rune = 0, 0
		if i >= 0 {
			x = str1[i] - '0'
		}
		if j >= 0 {
			y = str2[j] - '0'
		}
		bit := (x+y+carry)%rune(10) + '0'
		carry = (x + y + carry) / rune(10)
		ans = append(ans, bit)
	}

	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}
