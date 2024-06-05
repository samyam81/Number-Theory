package main

func countBeautifulPairs(nums []int) int {
    var count int =0
	n:=len(nums)

	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if AreCoPrime(firstdigit(nums[i]),lastdigit(nums[j])) {
					count++	
			}
		}
	}


	return count
}

func gcd(a, b int) int {
 	for a!=b{
		if a>b{
			a=a-b
		} else{
			b=b-a
		}
	}
	return a
}

func lastdigit(num int) int{
	return num%10
}

func firstdigit(num int) int{
	for num>=10{
		num/=10
	}
	return num
}

func AreCoPrime(x,y int) bool{
	return gcd(x,y)==1
}