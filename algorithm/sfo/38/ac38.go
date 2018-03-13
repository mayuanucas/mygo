package problem38

import "fmt"

func permutation(arrayChar []byte, start int) {
	if 1 >= len(arrayChar) {
		return
	}
	if start == len(arrayChar)-1 {
		for i := 0; i < len(arrayChar); i++ {
			fmt.Printf("%c", arrayChar[i])
		}
		fmt.Println()
	} else {
		for i := start; i < len(arrayChar); i++ {
			swap(arrayChar, start, i)
			permutation(arrayChar, start+1)
			swap(arrayChar, start, i)
		}
	}
}

func swap(array []byte, m, n int) {
	temp := array[m]
	array[m] = array[n]
	array[n] = temp
}
