package level1

// 基于数组实现栈
type StackV1 struct {
	data []interface{}
}

func (s *StackV1) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *StackV1) Pop() (out interface{}) {
	if s.Empty() {
		panic("stock empty")
	}
	out = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return out
}

func (s *StackV1) Peek() (out interface{}) {
	return s.data[len(s.data)-1]
}

func (s *StackV1) Empty() bool {
	return len(s.data) == 0
}
