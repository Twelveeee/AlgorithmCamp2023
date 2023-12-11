package level2

// 二进制求和
// https://leetcode.cn/problems/add-binary/description/

func addBinary(a string, b string) string {
	indexA, indexB := len(a)-1, len(b)-1
	ans := []byte{}

	if indexA == -1 {
		return b
	} else if indexB == -1 {
		return a
	}

	carry := 0
	for indexA >= 0 || indexB >= 0 {
		if indexA >= 0 {
			carry += int(a[indexA] - '0')
		}
		if indexB >= 0 {
			carry += int(b[indexB] - '0')
		}

		ans = append(ans, byte('0'+carry%2))
		carry = carry / 2

		indexA--
		indexB--
	}

	if carry > 0 {
		ans = append(ans, '1')
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return string(ans)
}
