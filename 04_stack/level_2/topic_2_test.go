package level2

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-1)
	obj.Push(0)
	obj.Push(-3)
	// fmt.Println(obj.GetMin())

	// obj.Push(9)
	// obj.Push(10)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
