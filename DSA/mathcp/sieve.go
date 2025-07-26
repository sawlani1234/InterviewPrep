package mathcp

import "fmt"

func Sieve() []bool {
	const N int = 10000

	primes := make([]bool, N)

	for i := 0; i < N; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= N; i++ {
		if primes[i] {
			for j := i * i; j < N; {
				primes[j] = false
				j = j + i
			}
		}
	}

	primes[0] = false
	primes[1] = false

	return primes
}

func PrintPrimesUptoN(n int) {
	primes := Sieve()

	for i := 0; i < n; i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}
}
