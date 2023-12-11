package level3

// https://leetcode.cn/problems/count-primes/
// 计算质数
func countPrimes(n int) int {
	primeList := make([]bool, n)
	ans := 0

	for i := 2; i < n; i++ {
		if !primeList[i] {
			ans += 1

			if i*i < n {
				for j := i * i; j < n; j += i {
					primeList[j] = true
				}
			}
		}
	}

	return ans
}
