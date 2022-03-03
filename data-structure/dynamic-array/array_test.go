package dynamic_array

import (
	"errors"
	"testing"
)

const (
	defaultCapacity = 16
)

type Array struct {
	data     []interface{}
	capacity int
	size     int
}

func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}

	return &Array{
		data:     make([]interface{}, capacity, capacity),
		capacity: capacity,
		size:     0,
	}
}

func NewDefault() *Array {
	return NewArray(defaultCapacity)
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) GetCapacity() int {
	return a.capacity
}

func (a *Array) Contains(target interface{}) bool {
	if a.IsEmpty() {
		return false
	}

	for _, d := range a.data {
		if d == target {
			return true
		}
	}

	return false
}

func (a *Array) checkIndex(index int) error {
	if index < 0 || index > a.GetSize() {
		return errors.New("index out of range error")
	}

	return nil
}

func (a *Array) resize(capacity int) {
	newArr := make([]interface{}, capacity, capacity)
	copy(newArr, a.data)
	a.data = newArr
	a.capacity = capacity
}

func (a *Array) Add(index int, target interface{}) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}

	if a.size >= a.capacity {
		a.resize(a.capacity << 1)
	}

	copy(a.data[index+1:], a.data[index:])
	a.data[index] = target
	a.size++

	return nil
}

func (a *Array) Append(target interface{}) error {
	return a.Add(a.size, target)
}

func (a *Array) Set(index int, target interface{}) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}

	a.data[index] = target
	return nil
}

func (a *Array) GetItem(index int) (interface{}, error) {
	if err := a.checkIndex(index); err != nil {
		return nil, err
	}

	return a.data[index], nil
}

func (a *Array) FindItem(target interface{}) (int, error) {
	for i, d := range a.data {
		if d == target {
			return i, nil
		}
	}

	return -1, errors.New("no item")
}

func (a *Array) Delete(target interface{}) bool {
	for i, d := range a.data {
		if d == target {
			copy(a.data[i:], a.data[i+1:])
			a.size--

			return true
		}
	}

	return false
}

func (a *Array) DeleteWithIndex(index int) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}

	copy(a.data[index:], a.data[index+1:])
	a.size--

	return nil
}

func TestArray(t *testing.T) {
	a := NewArray(2)
	t.Log(a)
	_ = a.Append(1)
	_ = a.Append(2)
	_ = a.Append(3)
	_ = a.Append(4)
	_ = a.Append(5)
	t.Log(a)

	err := a.Set(1, 22)
	t.Log(err, a)

	t.Log(a.GetItem(3))
	t.Log(a.FindItem(22))

	err = a.DeleteWithIndex(1)
	t.Log(err, a)

	b := a.Delete(3)
	t.Log(b, a)
}
