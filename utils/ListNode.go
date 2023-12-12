package utils

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
