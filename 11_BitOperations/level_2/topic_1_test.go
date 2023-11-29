package level2

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {

	var num1 uint32 = 0b00000000000000000000000000001011
	fmt.Println(hammingWeight(num1))

	var num2 uint32 = 0b00000000000000000000000010000000
	fmt.Println(hammingWeight(num2))

	var num3 uint32 = 0b11111111111111111111111111111101
	fmt.Println(hammingWeight(num3))
}
