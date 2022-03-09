package leetcode

import (
	"math"
	"testing"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/

func MaxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0)
	ans := make([]int, 0)

	for i := range nums {
		for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
			window = window[:len(window)-1]
		}

		for len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}

		window = append(window, i)

		if len(window) > 0 && i >= k-1 {
			ans = append(ans, nums[window[0]])
		}
	}

	return ans
}

// 执行超时
func MaxSlidingWindow1(nums []int, k int) []int {
	window := make([]int, k)
	ans := make([]int, 0)

	for i := k - 1; i < len(nums); i++ {
		window = nums[i-(k-1) : i+1]

		max := math.MinInt64
		for _, v := range window {
			if v > max {
				max = v
			}
		}

		ans = append(ans, max)
	}

	return ans
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	t.Log(MaxSlidingWindow(nums, k))

	nums = []int{1}
	k = 1
	t.Log(MaxSlidingWindow(nums, k))

	nums = []int{7, 2, 4}
	k = 2
	t.Log(MaxSlidingWindow(nums, k))
}
