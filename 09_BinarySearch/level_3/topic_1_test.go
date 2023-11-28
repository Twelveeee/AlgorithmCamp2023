package level3

import (
	// . "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	root1 := sortedArrayToBSTV2([]int{-10, -3, 0, 5, 9})
	fmt.Println(root1)
}
