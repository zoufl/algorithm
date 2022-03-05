package leetcode

import "testing"

// https://leetcode-cn.com/problems/implement-queue-using-stacks/

type MyQueue struct {
	input  []int
	output []int
}

func NewMyQueue() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

func (q *MyQueue) Pop() int {
	if len(q.output) == 0 {
		q.in2out()
	}

	res := q.output[len(q.output)-1]
	q.output = q.output[:len(q.output)-1]

	return res
}

func (q *MyQueue) Peek() int {
	if len(q.output) == 0 {
		q.in2out()
	}

	return q.output[len(q.output)-1]
}

func (q *MyQueue) in2out() {
	for len(q.input) > 0 {
		for i := len(q.input) - 1; i >= 0; i-- {
			q.output = append(q.output, q.input[i])
		}

		q.input = make([]int, 0)
	}
}

func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

func TestMyQueue(t *testing.T) {
	minStack := NewMyQueue()
	minStack.Push(1)
	minStack.Push(2)
	t.Log(minStack.Peek())
	t.Log(minStack.Pop())
	t.Log(minStack.Empty())
}
