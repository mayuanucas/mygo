package problem762

func countPrimeSetBits(L, R int) int {
	primes := make(map[int]bool, 0)
	primes[2] = true
	primes[3] = true
	primes[5] = true
	primes[7] = true
	primes[11] = true
	primes[13] = true
	primes[17] = true
	primes[19] = true

	cnt := 0
	for i := L; i <= R; i++ {
		bits := 0
		for n := i; n > 0; n /= 2 {
			bits += n & 1
		}
		if primes[bits] {
			cnt++
		}
	}
	return cnt
}
