package level2

// 反转字符串
// https://leetcode.cn/problems/reverse-string/description/

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
