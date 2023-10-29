package level2

import "math"

// 最小栈
// https://leetcode.cn/problems/min-stack/description/

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt64}}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	if val < s.minData[len(s.minData)-1] {
		s.minData = append(s.minData, val)
	} else {
		s.minData = append(s.minData, s.minData[len(s.minData)-1])
	}
}

func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.minData = s.minData[:len(s.minData)-1]
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	return s.minData[len(s.minData)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
