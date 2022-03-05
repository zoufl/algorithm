package leetcode

import (
	"math/rand"
	"testing"
)

// https://leetcode-cn.com/problems/shuffle-an-array/

type ShuffleArray struct {
	nums, base []int
}

func NewShuffleAnArray(nums []int) ShuffleArray {
	return ShuffleArray{
		nums: nums,
		base: append([]int{}, nums...),
	}
}

func (s *ShuffleArray) Reset() []int {
	copy(s.nums, s.base)

	return s.nums
}

// Fisher-Yates 洗牌算法
func (s *ShuffleArray) Shuffle() []int {
	n := len(s.nums)

	for i := range s.nums {
		j := i + rand.Intn(n-i)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}

	return s.nums
}

func TestShuffleAnArray(t *testing.T) {
	nums := []int{1, 2, 3}
	s := NewShuffleAnArray(nums)
	t.Log(s.Shuffle())
	t.Log(s.Reset())
	t.Log(s.Shuffle())
}
