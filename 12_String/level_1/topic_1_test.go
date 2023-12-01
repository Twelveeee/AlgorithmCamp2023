package level1

import (
	"fmt"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCaseV2("Hello"))
}
