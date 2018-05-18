package problem682

import "strconv"

func calPoints(ops []string) int {
	if nil == ops {
		return 0
	}

	sum := 0
	stack := make([]int, 0)
	for _, op := range ops {
		switch op {
		case "C":
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "D":
			val := stack[len(stack)-1] * 2
			sum += val
			stack = append(stack, val)
		case "+":
			val := stack[len(stack)-1] + stack[len(stack)-2]
			sum += val
			stack = append(stack, val)
		default:
			val, _ := strconv.Atoi(op)
			stack = append(stack, val)
			sum += val
		}
	}
	return sum
}
