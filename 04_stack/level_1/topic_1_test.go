package level1

import (
	"fmt"
	"testing"
)

func TestStackV1(t *testing.T) {
	data := make([]interface{}, 0, 20)
	stack := &StackV1{data}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
