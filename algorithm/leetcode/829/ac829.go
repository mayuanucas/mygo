package problem829

// 任何一个连续的序列可以写作-> (base+1,base+2,...,base+m)
// 对于本题，我们需要检查 (1+2+...+m)+ m*base==N 是否成立
// 即检查 sum=(1+2+...+m)   N-sum % m == 0
func consecutiveNumbersSum(N int) int {
	count := 0
	sum := 1
	m := 1

	for sum <= N {
		if 0 == (N-sum)%m {
			count++
		}
		m++
		sum = (1 + m) * m / 2
	}
	return count
}
