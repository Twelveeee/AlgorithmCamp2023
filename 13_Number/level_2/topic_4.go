package level2

// 判断 3的幂
// https://leetcode.cn/problems/power-of-three/description/

func isPowerOfThree(n int) bool {
	// n, i := 1, 0
	// for n < math.MaxInt32 {
	// 	n = int(math.Pow(3, float64(i)))
	// 	i++
	// 	fmt.Println(i, n)
	// }
	// return true

	return n > 0 && 1162261467%n == 0
}
