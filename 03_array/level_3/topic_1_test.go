package level1

import (
	"fmt"
	"testing"
)

func TestSingleNumberc(t *testing.T) {
	testCase := []struct {
		input  []int
		expect int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, v := range testCase {
		if v.expect != singleNumber(v.input) {
			fmt.Println(v)
		}
	}
	// isMonotonic()
}
