package level2

// 仅仅反转字母
// https://leetcode.cn/problems/reverse-only-letters/description/

func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	w := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > 'z' || s[i] < 'A' || (s[i] > 'Z' && s[i] < 'a') {
			continue
		}

		for w < len(ans) && (ans[w] > 'z' || ans[w] < 'A' || (ans[w] > 'Z' && ans[w] < 'a')) {
			w++
		}
		ans[w] = s[i]
		w++
	}
	return string(ans)
}
