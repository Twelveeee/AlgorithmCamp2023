package level3

import (
	"fmt"
	"testing"
)

func TestCalculatec(t *testing.T) {
	testCase := []struct {
		input  string
		expect int
	}{
		// {"3+2*2", 7},
		{" 3/2 ", 1},
		{" 3+5 / 2 ", 5},
	}

	for _, v := range testCase {
		if v.expect != calculate(v.input) {
			fmt.Println("error", calculate(v.input), v)
		}
	}
	// isMonotonic()
}
