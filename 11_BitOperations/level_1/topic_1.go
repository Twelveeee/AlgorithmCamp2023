package level1

import "fmt"

// 位运算和移位
func bitOperV1() {
	var a uint8 = 60 /* 60 = 0011 1100 */
	var b uint8 = 13 /* 13 = 0000 1101 */
	var c uint8 = 0

	var ab byte = 0b00111100 // 二进制
	var ax byte = 0x3C       // 十六进制
	var ao byte = 0o74       // 八进制

	fmt.Printf("a == ab: %v\n", a == ab)
	fmt.Printf("a == ax: %v\n", a == ax)
	fmt.Printf("a == ao: %v\n", a == ao)

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d \t二进制: %b\n", c, c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d \t二进制: %b\n", c, c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d \t二进制: %b\n", c, c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d \t二进制: %b\n", c, c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d \t二进制: %b\n", c, c)

}
