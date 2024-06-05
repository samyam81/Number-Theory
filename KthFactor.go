package main

func kthFactor(n int, k int) int {
	var factor[] int
    for i := 1; i <= n; i++ {
		if n%i==0 {
			factor = append(factor, i)
		}
	}
	if k>len(factor) {
		return -1
	}
	return factor[k-1]
}