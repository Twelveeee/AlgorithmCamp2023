package level2

// 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isLetter(s[left]) {
			left++
		}
		for left < right && !isLetter(s[right]) {
			right--
		}
		if !sameByte(s[left], s[right]) {
			return false
		}

		right--
		left++
	}

	return true
}

func isLetter(a byte) bool {
	// 48 57 65 90 97 122
	// '0', '9', 'A', 'Z', 'a', 'z'
	if a >= 'a' && a <= 'z' {
		return true
	}

	if a >= 'A' && a <= 'Z' {
		return true
	}

	if a >= '0' && a <= '9' {
		return true
	}

	return false
}

func sameByte(a, b byte) bool {
	if a == b {
		return true
	}

	if '0' >= a && a <= '9' {
		return false
	}

	if diff := int(a) - int(b); diff == -32 || diff == 32 {
		return true
	}

	return false
}
