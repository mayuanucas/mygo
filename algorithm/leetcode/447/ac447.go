package problem447

func numberOfBoomerangs(points [][]int) int {
	ans := 0

	for i := 0; i < len(points); i++ {
		dict := make(map[int]int, 0)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			d := getDistance(points[i], points[j])
			dict[d]++
		}

		for _, v := range dict {
			ans += v * (v - 1)
		}
	}

	return ans
}

func getDistance(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]

	return dx*dx + dy*dy
}
