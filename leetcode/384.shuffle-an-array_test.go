package leetcode

import (
	"math/rand"
	"testing"
)

// https://leetcode-cn.com/problems/shuffle-an-array/

type Solution struct {
	nums, base []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
		base: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.base)

	return s.nums
}

// Fisher-Yates 洗牌算法
func (s *Solution) Shuffle() []int {
	n := len(s.nums)

	for i := range s.nums {
		j := i + rand.Intn(n-i)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}

	return s.nums
}

func TestShuffleAnArray(t *testing.T) {
	nums := []int{1, 2, 3}
	s := Constructor(nums)
	t.Log(s.Shuffle())
	t.Log(s.Reset())
	t.Log(s.Shuffle())
}
