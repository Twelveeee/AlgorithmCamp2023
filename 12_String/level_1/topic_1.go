package level1

import (
	"strings"
)

// 转换成小写字母
// https://leetcode.cn/problems/to-lower-case/description/

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toLowerCaseV2(s string) string {
	ans := strings.Builder{}
	ans.Grow(len(s))
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}
		ans.WriteRune(ch)
	}
	return ans.String()
}
