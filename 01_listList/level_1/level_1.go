package level1

import (
	"fmt"
)

func level1() {
	fmt.Println("hola ")
}

func insertNode(node, newNode *LinkList, postion int) *LinkList {
	if postion < 0 {
		panic("error postion")
	}

	// 判断是不是节点
	if newNode.Next != nil {
		panic("newNode is a NodeList")
	}

	// 新建
	if node == nil {
		return newNode
	}

	// 头插
	if postion == 0 {
		newNode.Next = node
		node = newNode
		return node
	}

	head := node
	prev := node
	curr := node.Next

	for postion >= 0 {
		if postion == 1 {
			prev.Next = newNode
			newNode.Next = curr
			break
		}

		if curr == nil {
			panic("error out of length")
		}

		prev = curr
		curr = curr.Next
		postion--
	}

	return head
}

func delNode(node *LinkList, postion int) *LinkList {
	if postion < 0 {
		panic("error postion")
	}

	if node == nil {
		return node
	}

	// 头删
	if postion == 0 {
		return node.Next
	}

	prev := node
	curr := node.Next

	for postion >= 0 {
		if curr == nil {
			panic("error out of length")
		}
		if postion == 1 {
			prev.Next = curr.Next
			break
		}

		prev = curr
		curr = curr.Next
		postion--
	}

	return node
}
