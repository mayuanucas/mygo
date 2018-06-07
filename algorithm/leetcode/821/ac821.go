package problem821

import "math"

func shortestToChar(S string, C byte) []int {
	// 先获得 C 的各个坐标
	cId := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			cId = append(cId, i)
		}
	}

	ans := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			ans = append(ans, 0)
		} else {
			minDistance := getMinDistance(cId, i)
			ans = append(ans, minDistance)
		}
	}
	return ans
}

func getMinDistance(position []int, p int) int {
	d := math.MaxInt32
	for _, item := range position {
		temp := myAbs(p - item)
		if temp < d {
			d = temp
		}
	}
	return d
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
