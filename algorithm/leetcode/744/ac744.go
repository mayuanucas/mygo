package problem744

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			idx = i
			break
		}
	}
	return letters[idx]
}
