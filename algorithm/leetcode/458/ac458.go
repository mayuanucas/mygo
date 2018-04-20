package problem458

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0
	for mypow(minutesToTest/minutesToDie+1, pigs) < buckets {
		pigs++
	}
	return pigs
}

func mypow(a, b int) int {
	if 0 == b {
		return 1
	}
	if 1 == b {
		return a
	}
	if 0 == b&1 {
		return mypow(a*a, b/2)
	} else {
		return a * mypow(a*a, b/2)
	}
}
