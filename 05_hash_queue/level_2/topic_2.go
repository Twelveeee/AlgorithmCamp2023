package level2

// 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues/description/

type MyStack struct {
	outQueue, tempQueue []int
}

func ConstructorV2() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.tempQueue = append(s.tempQueue, x)
	for len(s.outQueue) > 0 {
		s.tempQueue = append(s.tempQueue, s.outQueue[0])
		s.outQueue = s.outQueue[1:]
	}

	s.tempQueue, s.outQueue = s.outQueue, s.tempQueue
}

func (s *MyStack) Pop() int {
	ans := s.outQueue[0]
	s.outQueue = s.outQueue[1:]
	return ans
}

func (s *MyStack) Top() int {
	return s.outQueue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.outQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
