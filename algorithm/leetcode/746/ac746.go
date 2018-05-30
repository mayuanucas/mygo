package problem746

func minCostClimbingStairs(cost []int) int {
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += myMin(cost[i+1], cost[i+2])
	}
	return myMin(cost[0], cost[1])
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
