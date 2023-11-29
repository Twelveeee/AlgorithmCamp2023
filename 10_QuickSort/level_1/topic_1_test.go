package level1

import (
	"fmt"
	"testing"
)

func TestQuickSortV1(t *testing.T) {
	var list1, list2, list3 []int

	list1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(quickSortV1(list1))
	list1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(quickSortV2(list1))

	list2 = []int{6, 5, 4, 3, 2, 1}
	fmt.Println(quickSortV1(list2))
	list2 = []int{6, 5, 4, 3, 2, 1}
	fmt.Println(quickSortV2(list2))

	list3 = []int{26, 53, 48, 15, 13, 48, 32, 15}
	fmt.Println(quickSortV1(list3))
	list3 = []int{26, 53, 48, 15, 13, 48, 32, 15}
	fmt.Println(quickSortV2(list3))
}
