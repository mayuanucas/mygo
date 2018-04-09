package problem342

// 4的幂应该满足的条件:
// 1. 大于0
// 2. 该数字的二进制表示中 仅有1个bit位数字为 1
// 3. 数字为 1 的bit位，应该出现在二进制数位的奇数位置
func isPowerOfFour(num int) bool {
	return num > 0 && (0 == num&(num-1)) && (num&0x55555555 != 0)
}
