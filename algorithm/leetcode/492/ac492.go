package problem492

import "math"

func constructRectangle(area int) []int {
	ans := make([]int, 0)

	num := int(math.Sqrt(float64(area)))
	for num > 0 {
		if area%num == 0 {
			if num >= area/num {
				ans = append(ans, num, area/num)
			} else {
				ans = append(ans, area/num, num)
			}
			break
		} else {
			num--
		}
	}

	return ans
}
