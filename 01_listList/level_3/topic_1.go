package level3

// 环形链表找入环节点
// https://leetcode.cn/problems/linked-list-cycle-ii/
func detectCycle(head *LinkNode) *LinkNode {
	fast, slow := head, head

	for {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			// 无环
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
