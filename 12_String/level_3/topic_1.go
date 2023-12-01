package level3

import (
	"math"
)

// https://leetcode.cn/problems/longest-common-prefix/description/
// 最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := math.MaxInt32
	var ans []byte

	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
		}
	}

	for i := 0; i < minLen; i++ {
		firstChar := strs[0][i]
		for _, str := range strs {
			if firstChar != str[i] {
				return string(ans)
			}
		}
		ans = append(ans, firstChar)
	}

	return string(ans)
}
