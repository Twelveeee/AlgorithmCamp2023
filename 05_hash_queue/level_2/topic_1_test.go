package level2

import (
	"fmt"
	"testing"
)

func TestMyQueue(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)

	fmt.Println(obj.Peek())

	// obj.Push(9)
	// obj.Push(10)

	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
