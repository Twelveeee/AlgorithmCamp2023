package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestIsPalindrome(t *testing.T) {
	headA := BuildListNode([]interface{}{1, 2, 2, 1})
	PrintListNode(headA)
	fmt.Println(isPalindrome(headA))
}
