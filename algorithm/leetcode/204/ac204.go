package problem204

func countPrimes(n int) int {
	notPrime := make([]bool, n)
	count := 0

	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++

			if temp := i * i; temp < n {
				for j := temp; j < n; j += i {
					notPrime[j] = true
				}
			}
		}
	}
	return count
}
