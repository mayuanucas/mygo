package problem137

func singleNumber(nums []int) int {
	dict := make(map[int]int, 0)

	for _, v := range nums {
		temp, ok := dict[v]
		if !ok {
			dict[v] = 1
		} else if 2 == temp {
			delete(dict, v)
		} else {
			dict[v]++
		}
	}
	for k := range dict {
		return k
	}

	return -1
}

func singleNumber2(nums []int) int {
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func singleNumber3(nums []int) int {
	one, two, three := 0, 0, 0
	for _, v := range nums {
		//当新来的为0时，two = two | 0，two不变，当新来的为1时，（当one此时为0，则two = two | 0，two不变；当one此时为1时，则two = two | 1，two变为1
		two |= one & v
		//新来的为0，one不变，新来为1时，one是0、1 交替改变
		one ^= v
		//当one和two为1是，three为1（此时代表要把one和two清零了）
		three = one & two
		//把one清0
		one &= ^three
		//把two清0
		two &= ^three
	}
	return one
}

// 遍历32次每次记录某位的出现的次数，如果不能被三整除，说明那个出现一次的就在该位有值，那么ans 或该位一下就可以了
func singleNumber4(nums []int) int {
	var ans = 0
	var i uint8
	for i = 0; i < 32; i++ {
		cnt := 0
		for _, v := range nums {
			if ((v >> i) & 1) == 1 {
				cnt++
			}
		}
		ans |= (cnt % 3) << i
	}
	return int(int32(ans))
}
