package level1

import (
	"fmt"
	"testing"
)

func TestLevel1(t *testing.T) {
	level1()
}

func TestInsertNode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	node1 := InitLinkList(1)
	node2 := InitLinkList(2)
	node3 := InitLinkList(3)
	node4 := InitLinkList(4)

	node1 = insertNode(node1, node2, 1)
	node1 = insertNode(node1, node3, 2)
	node1 = insertNode(node1, node4, 3)
	node1.PrintList()
}

func TestDelNode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	node1 := InitLinkList(1)
	node2 := InitLinkList(2)
	node3 := InitLinkList(3)
	node4 := InitLinkList(4)

	node1 = insertNode(node1, node2, 1)
	node1 = insertNode(node1, node3, 2)
	node1 = insertNode(node1, node4, 3)

	node1 = delNode(node1, 3)
	node1.PrintList()

}
