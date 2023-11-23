package level1

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestMergeTrees(t *testing.T) {
	node1 := BuildTree([]interface{}{1, 3, 2, 5})
	node2 := BuildTree([]interface{}{2, 1, 3, nil, 4, nil, 7})
	node3 := mergeTrees(node1, node2)

	fmt.Println(node3)
}

func TestMergeTreesV2(t *testing.T) {
	node1 := BuildTree([]interface{}{})
	node2 := BuildTree([]interface{}{1, 2})
	node3 := mergeTrees(node1, node2)

	fmt.Println(node3)
}
