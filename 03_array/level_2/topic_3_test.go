package level2

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5}
	// 5,0,1,2
	k := 1

	rotate(nums1, k)
	fmt.Println(nums1)
}
