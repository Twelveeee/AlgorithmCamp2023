package level2

// 用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/description/

type MyQueue struct {
	inStock, outStock []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStock = append(q.inStock, x)
}

func (q *MyQueue) Pop() int {
	if len(q.outStock) == 0 {
		q.tranceInStockToOutStock()
	}

	ans := q.outStock[len(q.outStock)-1]
	q.outStock = q.outStock[:len(q.outStock)-1]
	return ans
}

func (q *MyQueue) Peek() int {
	if len(q.outStock) == 0 {
		q.tranceInStockToOutStock()
	}

	return q.outStock[len(q.outStock)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStock) == 0 && len(q.outStock) == 0
}

func (q *MyQueue) tranceInStockToOutStock() {
	for len(q.inStock) > 0 {
		q.outStock = append(q.outStock, q.inStock[len(q.inStock)-1])
		q.inStock = q.inStock[:len(q.inStock)-1]
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
