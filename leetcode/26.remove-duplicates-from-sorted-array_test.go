package leetcode

import "testing"

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// 双指针
func RemoveDuplicates(nums []int) int {
	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	t.Log(RemoveDuplicates(nums), nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(RemoveDuplicates(nums), nums)
}
