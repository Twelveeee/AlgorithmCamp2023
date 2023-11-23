package level1

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	root1 := BuildTree([]interface{}{1, 2, 3, 4})
	root2 := BuildTree([]interface{}{1, 2, 3, 5})

	fmt.Print(isSameTree(root1, root2))
}
