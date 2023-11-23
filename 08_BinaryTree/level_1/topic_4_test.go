package level1

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestBinaryTreePaths(t *testing.T) {
	node1 := BuildTree([]interface{}{1, 2, 3, nil, 5})

	fmt.Println(binaryTreePaths(node1))

	node2 := BuildTree([]interface{}{1})
	fmt.Println(binaryTreePaths(node2))
}
