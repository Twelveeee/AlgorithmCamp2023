package level1

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestIsSymmetric(t *testing.T) {
	node1 := BuildTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(node1))
	node2 := BuildTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	fmt.Println(isSymmetric(node2))

}
