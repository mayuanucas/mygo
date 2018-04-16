package problem438

func findAnagrams(s string, p string) []int {
	dict := make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		dict[p[i]]++
	}

	ans := make([]int, 0)
	//two points, initialize count to p's length
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		//move right everytime, if the character exists in p's hash, decrease the count
		//current hash value >= 1 means the character is existing in p
		if dict[s[right]] >= 1 {
			count--
		}
		dict[s[right]]--
		right++

		//when the count is down to 0, means we found the right anagram
		//then add window's left to result list
		if 0 == count {
			ans = append(ans, left)
		}

		//if we find the window's size equals to p, then we have to move left (narrow the window) to find the new match window
		//++ to reset the hash because we kicked out the left
		//only increase the count if the character is in p
		//the count >= 0 indicate it was original in the hash, cuz it won't go below 0
		if right-left == len(p) {
			if v, ok := dict[s[left]]; ok && v >= 0 {
				count++
			}
			dict[s[left]]++
			left++
		}
	}

	return ans
}
