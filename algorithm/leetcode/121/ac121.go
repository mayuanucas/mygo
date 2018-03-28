package problem121

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		current := prices[i] - minPrice

		if current > maxProfit {
			maxProfit = current
		}
	}
	return maxProfit
}
