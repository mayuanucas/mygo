package problem443

import "strconv"

func compress(chars []byte) int {
	index, indexAns := 0, 0
	for index < len(chars) {
		currentChar := chars[index]
		count := 0
		for index < len(chars) && chars[index] == currentChar {
			index++
			count++
		}
		chars[indexAns] = currentChar
		indexAns++
		if 1 != count {
			for _, v := range []byte(strconv.Itoa(count)) {
				chars[indexAns] = v
				indexAns++
			}
		}
	}
	return indexAns
}
