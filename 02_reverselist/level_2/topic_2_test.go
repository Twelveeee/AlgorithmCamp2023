package level2

import (
	"fmt"
	"testing"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func TestSwapPairs(t *testing.T) {
	node := InitListNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = swapPairs(node)
	fmt.Println("**after**")

	node.PrintList()
}
