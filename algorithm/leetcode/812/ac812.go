package problem812

import "math"

// 从坐标计算三角形的面积公式为：
// area = |xa*yb + xb*yc + xc*ya - xa*yc - xb*ya - xc*yb| / 2
func largestTriangleArea(points [][]int) float64 {
	ans := 0.0
	for _, i := range points {
		for _, j := range points {
			for _, k := range points {
				ans = math.Max(ans, 0.5*math.Abs(float64(i[0]*j[1]+j[0]*k[1]+k[0]*i[1]-i[0]*k[1]-j[0]*i[1]-k[0]*j[1])))
			}
		}
	}
	return ans
}
