package level1

// 回文数
// https://leetcode.cn/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	num := 0
	// 12345

	for x > num {
		num = num*10 + x%10
		x = x / 10
	}

	return x == num || x == num/10

}
