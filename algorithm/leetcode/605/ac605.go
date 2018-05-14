package problem605

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if 0 == flowerbed[i] {
			prev, next := 0, 0
			if i == len(flowerbed)-1 {
				next = 0
			} else {
				next = flowerbed[i+1]
			}

			if 0 == i {
				prev = 0
			} else {
				prev = flowerbed[i-1]
			}

			if 0 == prev && next == 0 {
				flowerbed[i] = 1
				count++
			}
		}
	}
	// 花圃能够继续种植的花卉数目大于给的花卉数目，则成立
	return count >= n
}
