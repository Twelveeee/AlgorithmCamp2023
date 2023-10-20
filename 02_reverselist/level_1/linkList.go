package level1

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 打印链表
func (linkList *LinkNode) PrintList() {
	if linkList == nil {
		return
	}
	fmt.Println(linkList.Val)
	linkList.Next.PrintList()
}

func (linkList *LinkNode) ToList() []int {
	var res []int
	if linkList == nil {
		return res
	}

	res = append(res, linkList.Val)
	res = append(res, linkList.Next.ToList()...)
	return res
}

// 新建链表
func InitLinkNode(val int) *LinkNode {
	return &LinkNode{
		val, nil,
	}
}

// 新建链表
func InitLinkNodeByList(list []int) *LinkNode {
	tail := &LinkNode{
		0, nil,
	}
	head := tail

	for _, val := range list {
		tail.Next = &LinkNode{
			val, nil,
		}
		tail = tail.Next
	}

	return head.Next
}
