package level2

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums1 := []int{3, 2, 2, 3}
	val := 3

	// nums1 := []int{0, 0, 0}
	// val := 0

	val = removeElement(nums1, val)
	fmt.Println(val, nums1)
	// isMonotonic()
}
