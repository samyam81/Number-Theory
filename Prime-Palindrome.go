package main

import (
	"math"
)

func primePalindrome(n int) int {
	if n <= 2 {
		return 2
	}
	if n >= 9989900 {
		return 100030001
	}
	if n <= 11 && isPrime(n) {
		return n
	}

	for i := n + 1; i <= 20000000; i++ {
		if isPrime(i) && isPalindrome(i) {
			return i
		}
	}

	return -1
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(n int) bool {
	reversed, original := 0, n
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed == original
}
