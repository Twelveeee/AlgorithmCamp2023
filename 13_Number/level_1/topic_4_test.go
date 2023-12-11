package level1

import (
	"fmt"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	fmt.Println(convertToBase7(10))
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(101))
	fmt.Println(convertToBase7(-101))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(0))
}
