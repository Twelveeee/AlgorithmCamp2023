package level2

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	list1 := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(list1, 2))

	list2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(list2, 4))

	list3 := []int{0, 1}
	fmt.Println(findKthLargest(list3, 1))

	list4 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	fmt.Println(findKthLargest(list4, len(list4)))
}
