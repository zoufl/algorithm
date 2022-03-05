package stack

import (
	"container/list"
	"fmt"
	"testing"
)

type LinkedStack struct {
	data *list.List
	size int
}

func NewLinkedStack() *LinkedStack {
	s := &LinkedStack{
		data: list.New(),
		size: 0,
	}
	s.data.Init()

	return s
}

func (l *LinkedStack) Push(item interface{}) {
	l.data.PushBack(item)
	l.size++
}

func (l *LinkedStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}

	target := l.data.Back()
	l.data.Remove(target)
	l.size--

	return target.Value
}

func (l *LinkedStack) Peek() interface{} {
	if l.IsEmpty() {
		return nil
	}

	item := l.data.Back()

	return item.Value
}

func (l *LinkedStack) Clear() {
	l.data = list.New()
	l.size = 0
}

func (l *LinkedStack) Size() int {
	return l.size
}

func (l *LinkedStack) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedStack) Print() {
	if l.IsEmpty() {
		fmt.Println("{}")
	} else {
		fmt.Print("{")
		fmt.Print("size:", l.size, " ")
		fmt.Print("data:[")
		for n := l.data.Back(); n != nil; n = n.Prev() {
			fmt.Print(n.Value, " ")
		}
		fmt.Print("]")
		fmt.Println("}")
	}
}

func TestLinkedStack(t *testing.T) {
	s := NewLinkedStack()
	t.Log(s.IsEmpty())

	s.Push(11)
	s.Push(22)
	s.Push(33)
	t.Log(s.IsEmpty())
	t.Log(s.Size())

	s.Print()
	t.Log(s.Peek())
	t.Log(s.Pop())
	s.Print()

	s.Clear()
	s.Print()
}
