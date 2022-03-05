package stack

import (
	"testing"
)

type ArrayStack struct {
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0),
	}
}

func (a *ArrayStack) Push(item interface{}) {
	a.data = append(a.data, item)
}

func (a *ArrayStack) Pop() interface{} {
	if a.IsEmpty() {
		return nil
	}

	target := a.data[len(a.data)-1]
	a.data = a.data[:(len(a.data) - 1)]

	return target
}

func (a *ArrayStack) Peek() interface{} {
	if a.IsEmpty() {
		return nil
	}

	return a.data[len(a.data)-1]
}

func (a *ArrayStack) Clear() {
	a.data = make([]interface{}, 0)
}

func (a *ArrayStack) Size() int {
	return len(a.data)
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Size() == 0
}

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	t.Log(s.IsEmpty())

	s.Push(11)
	s.Push(22)
	s.Push(33)
	t.Log(s.IsEmpty())
	t.Log(s.Size())

	t.Log(s)
	t.Log(s.Peek())
	t.Log(s.Pop())
	t.Log(s)

	s.Clear()
	t.Log(s)
}
