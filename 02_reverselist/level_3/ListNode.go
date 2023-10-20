package level3

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
func (linkList *ListNode) PrintList() {
	if linkList == nil {
		return
	}
	fmt.Println(linkList.Val)
	linkList.Next.PrintList()
}

func (linkList *ListNode) ToList() []int {
	var res []int
	if linkList == nil {
		return res
	}

	res = append(res, linkList.Val)
	res = append(res, linkList.Next.ToList()...)
	return res
}

// 新建链表
func InitListNode(val int) *ListNode {
	return &ListNode{
		val, nil,
	}
}

// 新建链表
func InitListNodeByList(list []int) *ListNode {
	tail := &ListNode{
		0, nil,
	}
	head := tail

	for _, val := range list {
		tail.Next = &ListNode{
			val, nil,
		}
		tail = tail.Next
	}

	return head.Next
}
