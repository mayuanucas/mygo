package problem400

import "strconv"

func findNthDigit(n int) int {
	// 初始值定义 1 位数 共有 9 个, 起始数字为 1
	len, count, start := 1, 9, 1

	for n > len*count {
		n -= len * count
		len++
		count *= 10
		start *= 10
	}

	start += (n - 1) / len
	s := strconv.Itoa(start)
	return int(s[(n-1)%len] - '0')
}
