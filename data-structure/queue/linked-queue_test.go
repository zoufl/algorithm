package queue

import (
	"container/list"
	"testing"
)

type LinkedQueue struct {
	data *list.List
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		data: list.New(),
	}
}

// 进队列
func (q *LinkedQueue) Enqueue(item interface{}) {
	q.data.PushFront(item)
}

// 出队列
func (q *LinkedQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.data.Back()
	q.data.Remove(item)

	return item.Value
}

// 队列头元素
func (q *LinkedQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data.Back().Value
}

// 队列尾元素
func (q *LinkedQueue) Back() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data.Front().Value
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *LinkedQueue) Size() int {
	return q.data.Len()
}

func TestLinkedQueue(t *testing.T) {
	q := NewLinkedQueue()
	t.Log(q)
	t.Log(q.IsEmpty())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	t.Log(q.Size())
	t.Log(q)

	item := q.Front()
	t.Log(item, q)
	item = q.Back()
	t.Log(item, q)

	t.Log(q.Dequeue())
	t.Log(q.Dequeue())

	t.Log(q.Size())
}
