package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/majority-element/

// 哈希表
func MajorityElement2(nums []int) int {
	n := len(nums) / 2
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++

		if v, ok := m[num]; ok && v > n {
			return num
		}
	}

	return 0
}

// 排序取中间的值
func MajorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Boyer-Moore 投票算法
func MajorityElement(nums []int) int {
	count, candidate := 0, -1

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	t.Log(MajorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	t.Log(MajorityElement(nums))
}
