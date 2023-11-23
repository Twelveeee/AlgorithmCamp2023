package level2

import (
	"fmt"
	"testing"
)

func TestMaxDepthV3(t *testing.T) {
	node1 := &Node{1, nil}
	node2 := &Node{2, nil}
	node3 := &Node{3, nil}
	node4 := &Node{4, nil}
	node5 := &Node{5, nil}
	node6 := &Node{6, nil}

	node1.Children = append(node1.Children, node3, node2, node4)
	node3.Children = append(node3.Children, node5, node6)

	fmt.Println(maxDepthV3(node1))
}
