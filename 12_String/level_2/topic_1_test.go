package level2

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	list1 := []byte{'a', 'b', 'c', 'A', 'B', 'C'}
	fmt.Println(list1)
	reverseString(list1)
	fmt.Println(list1)

}
