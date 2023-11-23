package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestMinDepth(t *testing.T) {
	node1 := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})

	fmt.Println(minDepth(node1))

	node2 := BuildTree([]interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6})
	fmt.Println(minDepth(node2))

	node3 := BuildTree([]interface{}{})
	fmt.Println(minDepth(node3))
}
