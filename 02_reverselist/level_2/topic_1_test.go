package level2

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func TestReverseBetween(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4, 5})
	PrintListNode(node)

	node = reverseBetween(node, 2, 4)
	fmt.Println("**after**")

	PrintListNode(node)

}
