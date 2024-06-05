package main

func maximumPrimeDifference(nums []int) int {
   MaxDiff:=-1
   MinDiff:=-1
   for i := 0; i < len(nums); i++ {
		if isPrime(nums[i]) {
			if MinDiff==-1 {
				MinDiff=i
			}
			MaxDiff=i	
		}
   }
   if MinDiff!=-1 {
		return MaxDiff-MinDiff
   }
   return 0
}

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}
// 	sqrtN := int(math.Sqrt(float64(n)))
// 	for i := 3; i <= sqrtN; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }