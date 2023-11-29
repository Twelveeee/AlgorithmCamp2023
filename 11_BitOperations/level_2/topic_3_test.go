package level2

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	// var n1 uint32 = 0b00000010100101000001111010011100
	// var n1 uint32 = 1
	var n1 uint32 = 0b10000000000000000000000000000000
	fmt.Println(reverseBits(n1))
}
