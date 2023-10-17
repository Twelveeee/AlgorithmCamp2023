package level2

// 删除特定节点
// https://leetcode.cn/problems/remove-linked-list-elements/description/
func removeElements(head *LinkNode, val int) *LinkNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	tail := head
	pre := &LinkNode{}
	pre.Next = tail
	for tail != nil {
		if tail.Val == val {
			pre.Next = tail.Next
			tail = pre
		}

		pre = tail
		tail = tail.Next
	}

	return head
}
