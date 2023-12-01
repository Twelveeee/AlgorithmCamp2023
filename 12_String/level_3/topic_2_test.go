package level3

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	list1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(string(list1))
	len1 := compress(list1)
	fmt.Println(len1, string(list1[:len1]))
	fmt.Println()

	list3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(string(list3))
	len3 := compress(list3)
	fmt.Println(len3)
	fmt.Println(string(list3[:len3]))
	fmt.Println()

	list4 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(string(list4))
	len4 := compress(list4)
	fmt.Println(len4, string(list4[:len4]))
	fmt.Println()

	list5 := []byte{'a', 'b', 'c'}
	fmt.Println(string(list5))
	len5 := compress(list5)
	fmt.Println(len5, string(list5[:len5]))
	fmt.Println()

}
