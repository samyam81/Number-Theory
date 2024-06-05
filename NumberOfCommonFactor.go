package main
import "math"
func commonFactors(a int, b int) int {

    x := math.Max(float64(a), float64(b))

	var count int=0;

	for i := 1; i <= int(x); i++ {
		if a%i==0 && b%i==0 {
			count++
		}
	}
	return count
}