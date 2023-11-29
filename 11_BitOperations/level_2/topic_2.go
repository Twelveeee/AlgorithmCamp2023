package level2

// 比特位计数
// https://leetcode.cn/problems/counting-bits/

func countBits(n int) []int {
	dp := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		dp[i] = dp[i-highBit] + 1
	}

	return dp
}
