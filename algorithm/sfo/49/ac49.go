package problem49

func getUglyNumber(index int) int {
	if 0 >= index {
		return 0
	}

	uglyNumbers := make([]int, index)
	uglyNumbers[0] = 1
	nextUglyIndex := 1

	var pMultiply2, pMultiply3, pMultiply5 int
	for nextUglyIndex < index {
		uglyNumbers[nextUglyIndex] = min3(uglyNumbers[pMultiply2]*2, uglyNumbers[pMultiply3]*3, uglyNumbers[pMultiply5]*5)

		for uglyNumbers[pMultiply2]*2 <= uglyNumbers[nextUglyIndex] {
			pMultiply2++
		}
		for uglyNumbers[pMultiply3]*3 <= uglyNumbers[nextUglyIndex] {
			pMultiply3++
		}
		for uglyNumbers[pMultiply5]*5 <= uglyNumbers[nextUglyIndex] {
			pMultiply5++
		}
		nextUglyIndex++
	}
	return uglyNumbers[index-1]
}

func min3(num1, num2, num3 int) int {
	number := min(num1, num2)
	return min(num3, number)
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
