package leetcode

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-subarray/

// 动态规划
func MaxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func TestMaxSubArray(t *testing.T) {
	t.Log(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Log(MaxSubArray([]int{1}))
	t.Log(MaxSubArray([]int{5, 4, -1, 7, 8}))
	t.Log(MaxSubArray([]int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}))
}
