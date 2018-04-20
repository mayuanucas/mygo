package problem458

// 1. 一只猪在一小时内最多能验多少桶？
//
//一次喝一个桶的，15分钟后没挂再喝第二桶，一小时60分钟内可以喝 60/15 = 4 次，如果有5桶水，那个只要喝前4桶就只能第5桶是否有毒。
//
//因此一只小猪在一小时可以验5桶水
//
//
//2. 两只呢？
//
//既然一只能验5桶，那么用二维的思路，2只猪应该可以验5*5桶：
//
//猪A负责行，猪B负责列，每15分钟试喝一行/一列的所有5桶水，通过2只猪上天的时间能推断出毒水在几行几列。
//
//1   2   3   4   5
//
//6   7   8   9  10
//
//11 12 13 14 15
//
//16 17 18 19 20
//
//21 22 23 24 25
//
//
//3. N只
//
//如此类推到N只的情况，使用N维去分区，则5^N >= 1000即为解决本题的公式。

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
