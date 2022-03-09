package leetcode

import "testing"

// https://leetcode-cn.com/problems/design-circular-deque/

type MyCircularDeque struct {
	data     []int
	size     int
	capacity int
}

func NewMyCircularDeque(k int) *MyCircularDeque {
	return &MyCircularDeque{
		data:     make([]int, 0),
		capacity: k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	d.data = append([]int{value}, d.data...)
	d.size++

	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	d.data = append(d.data, value)
	d.size++

	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}

	d.data = d.data[1:]
	d.size--

	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}

	d.data = d.data[:len(d.data)-1]
	d.size--

	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}

	return d.data[0]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}

	return d.data[len(d.data)-1]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.capacity
}

func TestMyCircularDeque(t *testing.T) {
	deque := NewMyCircularDeque(3)
	t.Log(deque.InsertLast(1))
	t.Log(deque.InsertLast(2))
	t.Log(deque.InsertFront(3))
	t.Log(deque.InsertFront(4))
	t.Log(deque)
	t.Log(deque.GetRear())
	t.Log(deque.IsFull())
	t.Log(deque.DeleteLast())
	t.Log(deque)

	t.Log(deque.InsertFront(4))
	t.Log(deque.GetFront())
	t.Log(deque.GetRear())
	t.Log(deque)
}
