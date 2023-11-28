package level1

import (
	"fmt"
	"testing"
)

func TestBinarySearchV1(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearchV1(list1, 6))
}
