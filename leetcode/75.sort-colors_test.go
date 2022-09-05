package leetcode

import (
	"testing"
)

// https://leetcode.cn/problems/sort-colors/

// 双指针
func SortColors(nums []int) {
	p0, p2 := 0, len(nums)-1

	for i := 0; i <= p2; i++ {
		for i < p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}

		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// 单指针
func SortColors1(nums []int) {
	n := len(nums)

	if n == 0 {
		return
	}

	first := 0
	colors := []int{0, 1, 2}

	for i := 0; i < len(colors)-1; i++ {
		for j := first; j < n; j++ {
			if nums[j] == colors[i] {
				nums[first], nums[j] = nums[j], nums[first]
				first++
			}
		}
	}
}

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	SortColors(nums)
	t.Log(nums)

	nums = []int{2, 0, 1}
	SortColors(nums)
	t.Log(nums)
}
