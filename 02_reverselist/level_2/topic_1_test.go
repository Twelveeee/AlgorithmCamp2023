package level2

import (
	"fmt"
	"testing"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func TestReverseBetween(t *testing.T) {
	node := InitListNodeByList([]int{1, 2, 3, 4, 5})
	node.PrintList()

	node = reverseBetween(node, 2, 4)
	fmt.Println("**after**")

	node.PrintList()

}
