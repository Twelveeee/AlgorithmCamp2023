package level1

import (
	"fmt"
	"testing"
)

func TestStackV3(t *testing.T) {
	stack := InitStackV3()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
