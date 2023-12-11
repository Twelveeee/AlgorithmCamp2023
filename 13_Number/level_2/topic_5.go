package level2

// 判断 4的幂
// https://leetcode.cn/problems/power-of-four/description/

func isPowerOfFour(n int) bool {
	//      1	1
	//    100	4
	//  10000	16
	//1000000	64
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
