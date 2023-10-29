package level2

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	obj := ConstructorV2()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
	fmt.Println(obj.Pop())
}
