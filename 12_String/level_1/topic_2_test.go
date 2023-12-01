package level1

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("123"))
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("  -0012a42"))

}
