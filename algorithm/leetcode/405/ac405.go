package problem405

func toHex(num int) string {
	charMap := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	if 0 == num {
		return "0"
	}

	numUint32 := uint32(num)
	ans := make([]byte, 0)
	for 0 != numUint32 {
		ans = append(ans, charMap[numUint32&15])
		numUint32 >>= 4
	}

	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}
