package problem292

// Nim游戏: 每次最多取 1~n 个石头
// 精髓在于: 先手获胜的条件是 总石头数不能为 n+1 的倍数
func canWinNim(n int) bool {
	if n < 1 {
		return false
	}

	return 0 != n % 4
}
