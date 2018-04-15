package problem412

import "strconv"

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
