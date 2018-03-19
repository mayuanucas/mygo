package problem56

import (
	"unsafe"
	"errors"
)

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

func findNumberAppearOnce2(numbers []int) int {
	if nil == numbers || 0 >= len(numbers) {
		panic(errors.New("invalid input"))
	}

	bitSum := make([]int, 32)
	for i := 0; i < len(numbers); i++ {
		bitMask := 1
		for j := 31; j >= 0; j-- {
			bit := numbers[i] & bitMask
			if 0 != bit {
				bitSum[j] += 1
			}
			bitMask = bitMask << 1
		}
	}

	var result int
	for i := 0; i < 32; i++ {
		result = result << 1
		result += bitSum[i] % 3
	}
	return result
}
