package problem58

func reverse(str []rune) {
	copyStr := str

	begin, end := 0, len(copyStr)-1
	for begin < end {
		copyStr[begin], copyStr[end] = copyStr[end], copyStr[begin]
		begin++
		end--
	}
}

func reverseSentence(str string) string {
	if len(str) < 1 {
		return ""
	}

	copyCharArray := []rune(str)
	reverse(copyCharArray)

	var begin, end int
	for end <= len(copyCharArray) {
		if end == len(copyCharArray) {
			reverse(copyCharArray[begin:end])
			break
		}

		if copyCharArray[end] != ' ' {
			end++
		} else {
			reverse(copyCharArray[begin:end])
			end++
			begin = end
		}
	}
	return string(copyCharArray)
}
