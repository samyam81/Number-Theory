package main
import "math"

func findGCD(nums []int) int {
    
	max := math.MinInt64
	min := math.MaxInt64

	for _, number := range nums {
		if number>max {
				max=number			
		} 

		if number<min {
			min=number
		}
	}

	return GCD(max,min)
}

func GCD(a int, b int) int{
	for a!=b{
		if a>b{
			a=a-b
		} else{
			b=b-a
		}
	}
	return a
}