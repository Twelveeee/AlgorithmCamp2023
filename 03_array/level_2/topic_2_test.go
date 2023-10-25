package level2

import (
	"fmt"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	// nums1 := []int{3, 1, 2, 4}
	nums1 := []int{0, 1}

	nums1 = sortArrayByParity(nums1)
	fmt.Println(nums1)
}
