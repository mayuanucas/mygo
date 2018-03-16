package problem50

func firstNoRepeatingChar(str string) byte {
	if 0 >= len(str) {
		return 0
	}

	table := make([]int, 256)
	for i := 0; i < len(str); i++ {
		table[str[i]]++
	}
	for id, value := range table {
		if 1 == value {
			return byte(id)
		}
	}
	return 0
}
