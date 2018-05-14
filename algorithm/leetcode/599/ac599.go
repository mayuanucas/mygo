package problem599

func findRestaurant(list1 []string, list2 []string) []string {
	dict1 := make(map[string]int, len(list1))
	dict2 := make(map[string]int, len(list2))

	for id, v := range list1 {
		dict1[v] = id + 1
	}
	for id, v := range list2 {
		dict2[v] = id + 1
	}

	ans := make([]string, 0)
	min := 2001
	for k, v := range dict1 {
		if v2, ok := dict2[k]; ok && v+v2 <= min {
			if v+v2 < min {
				min = v + v2
				ans = ans[:0]
			}
			ans = append(ans, k)
		}
	}

	return ans
}
