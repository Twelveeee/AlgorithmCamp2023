package level2

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	headA := InitLinkNodeByList([]int{1, 2, 2, 1})
	headA.PrintList()
	fmt.Println(isPalindrome(headA))
}
