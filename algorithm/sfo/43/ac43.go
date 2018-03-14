package problem43

// 1~n 整数中 1 出现的次数 http://blog.csdn.net/yi_afly/article/details/52012593
func numberOf1Between1AndN(n int) int {
	ans := 0
	for i := 1; i <= n; i *= 10 {
		a := n / i
		b := n % i
		if 0 == a%10 {
			ans += i * a / 10
		} else if 1 == a%10 {
			ans += i*a/10 + b + 1
		} else {
			ans += (a/10 + 1) * i
		}
	}
	return ans
}
