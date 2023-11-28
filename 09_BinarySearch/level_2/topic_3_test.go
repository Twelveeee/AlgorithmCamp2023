package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestIsValidBST(t *testing.T) {
	node1 := BuildTree([]interface{}{2, 1, 3})
	fmt.Println(isValidBST(node1))

	node2 := BuildTree([]interface{}{1, nil, 2})
	fmt.Println(isValidBST(node2))

	node3 := BuildTree([]interface{}{})
	fmt.Println(isValidBST(node3))

	node4 := BuildTree([]interface{}{5, 1, 4, nil, nil, 3, 6})
	fmt.Println(isValidBST(node4))

	node5 := BuildTree([]interface{}{5, 4, 6, nil, nil, 3, 7})
	fmt.Println(isValidBST(node5))

	node6 := BuildTree([]interface{}{2, 2, 2})
	fmt.Println(isValidBST(node6))

}
