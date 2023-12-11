package level2

// 数组加一
// https://leetcode.cn/problems/plus-one/description/

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if i == 0 && digits[i] == 9 {
			digits = append(digits, 0)
			digits[0] = 1
			break
		}

		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}

	return digits
}
