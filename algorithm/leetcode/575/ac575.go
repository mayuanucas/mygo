package problem575

func distributeCandies(candies []int) int {
	mySet := make(map[int]int, 0)

	for _, v := range candies {
		mySet[v]++
	}

	if len(mySet) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(mySet)
}
