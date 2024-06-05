package main

func isThree(n int) bool {
	if n < 4 {
		return false
	}

	var count uint = 0

	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			count++
		}
		if count > 3 {
			return false
		}
	}

	return count == 2
}
