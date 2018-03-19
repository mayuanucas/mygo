package problem57

import "fmt"

func findNumbers(data []int, sum int) (int, int, bool) {
	var found bool
	if nil == data || 1 > len(data) {
		return 0, 0, found
	}

	ahead, behind := 0, len(data)-1
	for ahead < behind {
		currentSum := data[ahead] + data[behind]
		if currentSum == sum {
			found = true
			return data[ahead], data[behind], found
		} else if currentSum < sum {
			ahead++
		} else {
			behind--
		}
	}
	return 0, 0, found
}

func findContinuousSequence(sum int) {
	if 3 > sum {
		return
	}

	middle := (sum + 1) / 2
	small, big := 1, 2
	currentSum := small + big
	for small < middle {
		if currentSum == sum {
			printContinuousSequence(small, big)
		}

		for currentSum > sum && small < middle {
			currentSum -= small
			small++

			if currentSum == sum {
				printContinuousSequence(small, big)
			}
		}
		big++
		currentSum += big
	}
}

func printContinuousSequence(small, big int) {
	for i := small; i <= big; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
