package level2

// 反转字符串 II
// https://leetcode.cn/problems/reverse-string-ii/

// abcd efgh   2
// abcd efgh i 2
func reverseStr(s string, k int) string {
	ans := []byte(s)
	times := len(s) / (2 * k)
	if len(s)%(2*k) != 0 {
		times++
	}

	for i := 0; i <= times; i++ {
		left, rihgt := i*2*k, (i*2*k)+k
		if rihgt > len(s) {
			rihgt = len(s)
		}

		for j := 0; j < (rihgt-left)/2; j++ {
			ans[left+j], ans[rihgt-j-1] = ans[rihgt-j-1], ans[left+j]
		}
	}

	return string(ans)
}
