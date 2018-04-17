package problem409

//1、统计所有字母的出现频率（分大小写）
//2、统计只出现奇数次数字母的个数
//3、如果2中结果不为0，字符串的长度减去2中的字母个数+1
//
//其中3的意思是，保留出现次数最多的那个奇数字母，剩下的需要全部减1变成偶数去构造
func longestPalindrome(s string) int {
	dict := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	odd := 0
	for _, v := range dict {
		if v&1 != 0 {
			odd++
		}
	}

	ans := len(s)
	if 0 != odd {
		ans -= odd - 1
	}
	return ans
}
