package level2

import (
	"fmt"
	"testing"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func TestAddTwoNumbers(t *testing.T) {
	node := InitListNodeByList([]int{7, 2, 4, 3})
	node2 := InitListNodeByList([]int{5, 6, 4})

	node = addTwoNumbers(node, node2)
	fmt.Println("**after**")

	node.PrintList()
}
