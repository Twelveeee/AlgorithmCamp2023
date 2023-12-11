package level3

// 辗转相除法(欧几里得算法)
// 求最大公约数 GCD(Greatest Common Divisor)
// eg1:
// 12，18的公因数有：1，2，3，6。
// gcd（12，18）=gcd(18,12%18)=gcd(18,12)
// gcd（18，12）=gcd(12,18%12)=gcd(12,6)
// gcd（12，6）=gcd(6,12%6)=gcd(6,0)
// return 6
func gcd(a, b int) int {
	k := a % b
	a, b = b, k
	for k != 0 {
		k = a % b
		a, b = b, k
	}
	return a
}
