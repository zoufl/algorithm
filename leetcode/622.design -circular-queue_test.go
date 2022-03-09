package leetcode

import "testing"

// https://leetcode-cn.com/problems/design-circular-queue/

type MyCircularQueue struct {
	data     []int
	size     int
	capacity int
	head     int
	tail     int
}

func NewMyCircularQueue(k int) *MyCircularQueue {
	return &MyCircularQueue{
		data:     make([]int, k, k),
		capacity: k,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.data[q.tail] = value
	q.tail = (q.tail + 1) % len(q.data)
	q.size++

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.head = (q.head + 1) % len(q.data)
	q.size--

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	tail := q.tail - 1
	if tail < 0 {
		tail = q.capacity - 1
	}

	return q.data[tail]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.size == q.capacity
}

func TestMyCircularQueue(t *testing.T) {
	q := NewMyCircularQueue(3)
	t.Log(q.EnQueue(1))
	t.Log(q.EnQueue(2))
	t.Log(q.EnQueue(3))
	t.Log(q.EnQueue(4))

	t.Log(q)

	t.Log(q.Front())
	t.Log(q.Rear())
	t.Log(q.IsFull())
	t.Log("dequeue", q.DeQueue())
	t.Log(q.EnQueue(4))
	t.Log(q.Front())
	t.Log(q.Rear())

	t.Log(q)
}
