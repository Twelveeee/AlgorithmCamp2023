package level2

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	list1 := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(list1))

	list2 := []int{2, 1}
	fmt.Println(findMin(list2))
}
