package level1

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
