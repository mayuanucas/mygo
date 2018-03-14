package problem42

func findGreatestSumOfSubArray(data []int) (int, bool) {
	if nil == data || len(data) <= 0 {
		return 0, false
	}
	var currentSum, greatestSum int
	for i := 0; i < len(data); i++ {
		if currentSum <= 0 {
			currentSum = data[i]
		} else {
			currentSum += data[i]
		}
		if currentSum > greatestSum {
			greatestSum = currentSum
		}
	}
	return greatestSum, true
}
