package level3

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	testCase := []struct {
		input  []string
		expect int
	}{
		// {[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, v := range testCase {
		if v.expect != evalRPN(v.input) {
			fmt.Println("error", evalRPN(v.input), v)
		}
	}
}
