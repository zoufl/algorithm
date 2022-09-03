package leetcode

import "testing"

// https://leetcode.cn/problems/jump-game/

func CanJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}

		k = MaxJump(k, i+nums[i])
	}

	return true
}

// 贪心
func CanJump1(nums []int) bool {
	n, rightmost := len(nums), 0

	for i := 0; i < len(nums); i++ {
		if i <= rightmost {
			rightmost = MaxJump(rightmost, i+nums[i])

			if rightmost >= n-1 {
				return true
			}
		}
	}

	return false
}

func MaxJump(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestCanJump(t *testing.T) {
	t.Log(CanJump([]int{2, 3, 1, 1, 4}))
	t.Log(CanJump([]int{3, 2, 1, 0, 4}))
}
