package level2

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node1 := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})

	fmt.Println(maxDepth(node1))

	node2 := BuildTree([]interface{}{1, nil, 2})
	fmt.Println(maxDepth(node2))

	node3 := BuildTree([]interface{}{})
	fmt.Println(maxDepth(node3))
}
