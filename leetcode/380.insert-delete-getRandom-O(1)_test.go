package leetcode

import (
	"math/rand"
	"testing"
)

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1/

type RandomizedSet struct {
	m    map[int]int
	data []int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		m:    map[int]int{},
		data: make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; ok {
		return false
	}

	s.data = append(s.data, val)
	s.m[val] = len(s.data) - 1

	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if i, ok := s.m[val]; ok {
		lastIndex := len(s.data) - 1
		s.m[s.data[lastIndex]] = i
		s.data[i], s.data[lastIndex] = s.data[lastIndex], s.data[i]
		s.data = s.data[:lastIndex]
		delete(s.m, val)

		return true
	}

	return false
}

func (s *RandomizedSet) GetRandom() int {
	return s.data[rand.Intn(len(s.data))]
}

func TestRandomizedSet(t *testing.T) {
	randomizedSet := NewRandomizedSet()
	t.Log(randomizedSet.Insert(1))
	t.Log(randomizedSet)
	t.Log(randomizedSet.Remove(2))
	t.Log(randomizedSet)
	t.Log(randomizedSet.Insert(2))
	t.Log(randomizedSet)
	t.Log(randomizedSet.GetRandom())
	t.Log(randomizedSet.Remove(1))
	t.Log(randomizedSet)
	t.Log(randomizedSet.Insert(2))
	t.Log(randomizedSet)
	t.Log(randomizedSet.GetRandom())

	//t.Log(randomizedSet.Insert(0))
	//t.Log(randomizedSet)
	//t.Log(randomizedSet.Insert(1))
	//t.Log(randomizedSet)
	//t.Log(randomizedSet.Remove(0))
	//t.Log(randomizedSet)
	//t.Log(randomizedSet.Insert(2))
	//t.Log(randomizedSet)
	//t.Log(randomizedSet.Remove(1))
	//t.Log(randomizedSet)
	//t.Log(randomizedSet.GetRandom())
}
