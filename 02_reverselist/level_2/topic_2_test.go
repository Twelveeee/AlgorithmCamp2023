package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func TestSwapPairs(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = swapPairs(node)
	fmt.Println("**after**")

	PrintListNode(node)
}
