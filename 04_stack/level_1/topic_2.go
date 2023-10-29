package level1

// 基于链表实现栈
type StackV2 struct {
	top *Node
}

type Node struct {
	Pre *Node
	Val interface{}
}

func (s *StackV2) Push(v interface{}) {
	s.top = &Node{s.top, v}
}

func (s *StackV2) Pop() (out interface{}) {
	if s.Empty() {
		panic("stock empty")
	}
	outNode := s.top
	s.top = outNode.Pre

	return outNode.Val
}

func (s *StackV2) Peek() (out interface{}) {
	return s.top.Val
}

func (s *StackV2) Empty() bool {
	return s.top == nil
}
