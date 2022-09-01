package leetcode

import "testing"

//https://leetcode.cn/problems/next-permutation/

// 两遍扫描
func NextPermutation(nums []int) {
	n := len(nums)
	i, j := n-2, n-1

	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		k := n - 1
		for k >= 0 && nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	NextPermutation(nums)
	t.Log(nums)

	nums = []int{3, 2, 1}
	NextPermutation(nums)
	t.Log(nums)

	nums = []int{1, 1, 5}
	NextPermutation(nums)
	t.Log(nums)
}
