package level3

import (
	"fmt"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	obj := Constructor()
	obj.AddNum(-1)
	obj.AddNum(-2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-3)
	fmt.Println(obj.FindMedian())
	// obj.AddNum(-4)
	// obj.AddNum(-3)
	// obj.AddNum(-4)

	fmt.Println(obj.maxHeap)
	fmt.Println(obj.minHeap)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-4)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(-5)
	// fmt.Println(obj.FindMedian())
}
