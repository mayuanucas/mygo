package problem56

import "unsafe"

func findNumbersAppearOnce(data []int) (int, int) {
	if nil == data || len(data) < 2 {
		return 0, 0
	}

	var result int
	for i := 0; i < len(data); i++ {
		result ^= data[i]
	}
	indexOf1 := findFirstBitIs1(result)
	var num1, num2 int
	for j := 0; j < len(data); j++ {
		if isBit1(data[j], indexOf1) {
			num1 ^= data[j]
		} else {
			num2 ^= data[j]
		}
	}
	return num1, num2
}

func findFirstBitIs1(num int) uint {
	var indexBit uint
	for (0 == num&1) && (indexBit < uint(unsafe.Sizeof(int(0))*8)) {
		num = num >> 1
		indexBit++
	}
	return indexBit
}

func isBit1(num int, indexBit uint) bool {
	num = num >> indexBit
	if 1 == num&1 {
		return true
	}
	return false
}
