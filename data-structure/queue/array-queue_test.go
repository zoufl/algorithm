package queue

import "testing"

type ArrayQueue struct {
	data []interface{}
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		data: make([]interface{}, 0),
	}
}

// 入队
func (q *ArrayQueue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

// 出队
func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]

	return item
}

// 队列头元素
func (q *ArrayQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data[0]
}

// 队列尾元素
func (q *ArrayQueue) Back() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data[len(q.data)-1]
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.data)
}

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue()
	t.Log(q)
	t.Log(q.IsEmpty())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
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
