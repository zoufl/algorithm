package leetcode

import (
	"testing"
)

type MyStackUsingQueue struct {
	input  []int
	output []int
}

func NewStackUsingQueue() *MyStackUsingQueue {
	return &MyStackUsingQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

func (s *MyStackUsingQueue) Push(x int) {
	s.input = append(s.input, x)
}

func (s *MyStackUsingQueue) Pop() int {
	s.in2out()

	item := s.output[len(s.output)-1]
	s.output = s.output[:len(s.output)-1]

	return item
}

func (s *MyStackUsingQueue) Top() int {
	s.in2out()

	return s.output[len(s.output)-1]
}

func (s *MyStackUsingQueue) in2out() {
	if len(s.input) > 0 {
		for _, item := range s.input {
			s.output = append(s.output, item)
		}

		s.input = []int{}
	}
}

func (s *MyStackUsingQueue) Empty() bool {
	return len(s.input) == 0 && len(s.output) == 0
}

func TestMyStackUsingQueue(t *testing.T) {
	s := NewStackUsingQueue()
	s.Push(1)
	s.Push(2)
	t.Log(s)
	t.Log(s.Top())
	t.Log(s.Pop())
	t.Log(s.Empty())
	s.Push(3)
	t.Log(s.Top())
	t.Log(s)
}
