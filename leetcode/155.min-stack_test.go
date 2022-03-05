package leetcode

import (
	"math"
	"testing"
)

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	data []int
	min  int64
}

func NewMinStack() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  math.MaxInt64,
	}
}

func (m *MinStack) Push(val int) {
	m.data = append(m.data, val)

	if int64(val) < m.min {
		m.min = int64(val)
	}
}

func (m *MinStack) Pop() {
	top := m.Top()
	m.data = m.data[:len(m.data)-1]

	if len(m.data) != 0 && m.min == int64(top) {
		m.min = int64(m.data[0])
		for _, d := range m.data {
			if int64(d) < m.min {
				m.min = int64(d)
			}
		}
	}

	if len(m.data) == 0 {
		m.min = math.MaxInt64
	}
}

func (m *MinStack) Top() int {
	return m.data[len(m.data)-1]
}

func (m *MinStack) GetMin() int {
	return int(m.min)
}

func TestMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	minStack.GetMin()
}
