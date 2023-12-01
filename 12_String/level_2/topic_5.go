package level2

// 字符串中的第一个唯一字符
// https://leetcode.cn/problems/valid-palindrome/description/

func firstUniqChar(s string) int {
	list := make([]int, 26)
	for _, v := range s {
		list[v-'a']++
	}

	for i, v := range s {
		if list[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
