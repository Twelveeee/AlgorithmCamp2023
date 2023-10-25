package level1

import (
	"fmt"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	testCase := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{1, 3, 2}, false},
	}

	for _, v := range testCase {
		if v.expect != isMonotonicV2(v.input) {
			fmt.Println(v)
		}
	}
	// isMonotonic()
}
