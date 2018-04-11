package problem389

func findTheDifference(s string, t string) byte {
	dict1 := make(map[rune]int, 0)
	for _, v := range s {
		dict1[v]++
	}

	dict2 := make(map[rune]int, 0)
	for _, v := range t {
		dict2[v]++
	}

	for k2, v2 := range dict2 {
		if v1 := dict1[k2]; v1 != v2 {
			return byte(k2)
		}
	}
	return byte(0)
}
