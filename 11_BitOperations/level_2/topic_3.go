package level2

// 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/description/

// 将整数 n 向右移动一位并与 m1 进行按位与操作，然后将结果与 n 向左移动一位并与 m1 进行按位与操作的结果进行按位或操作。这一步交换了相邻的位。
// 将上一步的结果向右移动两位并与 m2 进行按位与操作，然后将结果与上一步的结果向左移动两位并与 m2 进行按位与操作的结果进行按位或操作。这一步交换了每两位的块。
// 将上一步的结果向右移动四位并与 m4 进行按位与操作，然后将结果与上一步的结果向左移动四位并与 m4 进行按位与操作的结果进行按位或操作。这一步交换了每四位的块。
// 将上一步的结果向右移动八位并与 m8 进行按位与操作，然后将结果与上一步的结果向左移动八位并与 m8 进行按位与操作的结果进行按位或操作。这一步交换了每八位的块。
// 将上一步的结果向右移动16位，并将结果与上一步的结果向左移动16位的结果进行按位或操作，完成整个整数的二进制位颠倒。
const (
	m1 = 0b01010101010101010101010101010101
	m2 = 0b00110011001100110011001100110011
	m4 = 0b00001111000011110000111100001111
	m8 = 0b00000000111111110000000011111111
)

func reverseBits(num uint32) uint32 {
	num = num>>1&m1 | num&m1<<1
	num = num>>2&m2 | num&m2<<2
	num = num>>4&m4 | num&m4<<4
	num = num>>8&m8 | num&m8<<8
	return num>>16 | num<<16
}
