package level1

// 建立虚拟头节点辅助翻转
func reverseListV1(head *LinkNode) *LinkNode {

	dummyHead, temp := &LinkNode{}, &LinkNode{}
	dummyHead.Next = head
	pre, curr := dummyHead, dummyHead.Next
	for curr.Next != nil {
		temp = curr.Next
		curr.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyHead.Next
}

// 直接操作链表实现翻转
func reverseListV2(head *LinkNode) *LinkNode {

	curr := head
	temp := &LinkNode{}
	var pre *LinkNode

	for curr != nil {
		temp = curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}

	return pre
}

// 递归
func reverseListV3(head *LinkNode) *LinkNode {

	if head.Next == nil || head == nil {
		return head
	}

	// 尾节点
	tailHead := reverseListV3(head.Next)

	// 将当前节点的下一个节点的 Next 指针指向当前节点，实现翻转
	head.Next.Next = head
	head.Next = nil

	return tailHead
}
