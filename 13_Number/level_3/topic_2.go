package level3

import "math"

// 判断是不是质数
func isPrime(num int) bool {
	max := math.Sqrt(float64(num))
	for i := 2; i <= int(max); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
