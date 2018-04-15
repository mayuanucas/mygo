package problem412

import (
	"strconv"
	"fmt"
)

func fizzBuzz(n int) []string {
	ans := make([]string, 0)

	for i := 1; i <= n; i++ {
		ans = append(ans, number(i))
	}
	return ans
}

func number(n int) string {
	if 0 == n%3 && 0 == n%5 {
		return "FizzBuzz"
	} else if 0 == n%5 {
		return "Buzz"
	} else if 0 == n%3 {
		return "Fizz"
	} else {
		return strconv.Itoa(n)
	}
}

func fizzBuzz2(n int) []string {
	ans := make([]string, 0)

	for i, fizz, buzz := 1, 0, 0; i <= n; i++ {
		fizz++
		buzz++
		if 3 == fizz && 5 == buzz {
			ans = append(ans, "FizzBuzz")
			fizz = 0
			buzz = 0
		} else if 3 == fizz {
			ans = append(ans, "Fizz")
			fizz = 0
		} else if 5 == buzz {
			ans = append(ans, "Buzz")
			buzz = 0
		} else {
			ans = append(ans, fmt.Sprintf("%d", i))
		}
	}
	return ans
}
