package level1

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	nums1 := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums1)
	fmt.Println(nums1)
}
