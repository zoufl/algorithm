package queue

import (
	"errors"
	"testing"
)

type CircularQueue struct {
	data     []interface{}
	capacity int
	size     int
	head     int
	tail     int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:     make([]interface{}, capacity, capacity),
		capacity: capacity,
	}
}

// 入队
func (q *CircularQueue) Enqueue(item interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.tail] = item
	q.tail = (q.tail + 1) % q.Capacity()
	q.size++

	return nil
}

//出队
func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.data[q.head]
	q.head = (q.head + 1) % q.Capacity()
	q.size--

	return item
}

//队列头元素
func (q *CircularQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data[q.head]
}

//队列尾部元素
func (q *CircularQueue) Rear() interface{} {
	if q.IsEmpty() {
		return nil
	}

	tail := q.tail - 1
	if tail < 0 {
		tail = q.capacity - 1
	}

	return q.data[tail]
}

// 队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.Size() == 0
}

// 队列是否已满
func (q *CircularQueue) IsFull() bool {
	return q.Size() == q.Capacity()
}

func (q *CircularQueue) Size() int {
	return q.size
}

func (q *CircularQueue) Capacity() int {
	return q.capacity
}

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(3)

	t.Log(q.Enqueue(1))
	t.Log(q.Enqueue(2))
	t.Log(q.Enqueue(3))
	t.Log(q.Enqueue(4))

	t.Log("size: ", q.Size(), q)

	t.Log(q.Front())
	t.Log(q.Rear())
	t.Log(q.IsFull())
	t.Log("dequeue", q.Dequeue())
	t.Log("dequeue", q.Dequeue())
	t.Log("dequeue", q.Dequeue())
	t.Log("dequeue", q.Dequeue())

	t.Log("size: ", q.Size(), q)

	t.Log(q.Enqueue(5))
	t.Log("size: ", q.Size(), q)
	t.Log(q.Front())
	t.Log(q.Rear())

}
