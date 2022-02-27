package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/single-number/

// 位运算（线性时间复杂度，不使用额外空间）
func SingleNumber(nums []int) int {
	single := 0

	for _, num := range nums {
		single ^= num
	}

	return single
}

// 排序之后遍历查找
func SingleNumber1(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	t.Log(SingleNumber(nums))

	nums = []int{4, 1, 2, 1, 2}
	t.Log(SingleNumber(nums))
}
