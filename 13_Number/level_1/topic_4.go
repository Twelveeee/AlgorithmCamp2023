package level1

// 七进制数
// https://leetcode.cn/problems/base-7/description/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var ans = []byte{}

	flag := false
	if num < 0 {
		num *= -1
		flag = true
	}

	for num > 0 {
		ans = append(ans, byte('0'+num%7))
		num = num / 7
	}

	if flag {
		ans = append(ans, '-')
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return string(ans)
}
