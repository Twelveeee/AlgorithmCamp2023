package level2

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	node1 := BuildListNode([]interface{}{1, 4, 5})
	node2 := BuildListNode([]interface{}{1, 3, 4})
	node3 := BuildListNode([]interface{}{2, 6})

	nodeList := make([]*ListNode, 0)
	nodeList = append(nodeList, node1)
	nodeList = append(nodeList, node2)
	nodeList = append(nodeList, node3)

	node := mergeKLists(nodeList)
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}
