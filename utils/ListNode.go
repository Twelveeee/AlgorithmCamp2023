package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(list []interface{}) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for _, v := range list {
		cur.Next = &ListNode{Val: v.(int)}
		cur = cur.Next
	}

	return head.Next
}

func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	PrintListNode(head.Next)
}

func InitLinkNode(val int) *ListNode {
	return &ListNode{
		val, nil,
	}
}
