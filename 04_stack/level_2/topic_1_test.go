package level2

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCase := []struct {
		input  string
		expect bool
	}{
		{"()", true},
		// {"()[]{}", true},
		// {"(]", false},
	}

	for _, v := range testCase {
		if v.expect != isValid(v.input) {
			fmt.Println("error", v)
		}
	}
}
