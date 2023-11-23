package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestIsBalanced(t *testing.T) {
	node1 := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})

	fmt.Println(isBalanced(node1))

	node2 := BuildTree([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})
	fmt.Println(isBalanced(node2))

	node3 := BuildTree([]interface{}{})
	fmt.Println(isBalanced(node3))
}
