package level3

// https://leetcode.cn/problems/string-compression/
// 压缩字符串
func compress(chars []byte) int {
	if len(chars) <= 1 {
		return 1
	}

	pre, write := 0, 0
	for i := 0; i < len(chars); i++ {
		if i == len(chars)-1 || chars[pre] != chars[i+1] {
			chars[write] = chars[pre]
			write++
			count := i - pre + 1
			if count > 1 {
				tempWrite := write
				for count > 0 {
					chars[write] = '0' + byte(count%10)
					count /= 10
					write++
				}
				for j := 0; j < (write-tempWrite)/2; j++ {
					chars[tempWrite+j], chars[write-j-1] = chars[write-j-1], chars[tempWrite+j]
				}
			}
			pre = i + 1
		}
	}

	return write
}
