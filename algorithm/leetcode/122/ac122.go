package problem122

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
