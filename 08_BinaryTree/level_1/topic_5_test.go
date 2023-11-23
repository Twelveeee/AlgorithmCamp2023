package level1

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestHasPathSum(t *testing.T) {
	node1 := BuildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1})

	fmt.Println(hasPathSum(node1, 22))

	node2 := BuildTree([]interface{}{1, 2, 3})
	fmt.Println(hasPathSum(node2, 5))

	node3 := BuildTree([]interface{}{})
	fmt.Println(hasPathSum(node3, 0))
}
