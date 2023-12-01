package level2

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("  "))
	fmt.Println(isPalindrome("0P"))
	// fmt.Println('0', '9', 'A', 'Z', 'a', 'z')

}
