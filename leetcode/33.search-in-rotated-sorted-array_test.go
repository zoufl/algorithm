package leetcode

import "testing"

// https://leetcode.cn/problems/search-in-rotated-sorted-array/

// 二分查找
func SearchInRotatedSortedArray(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	t.Log(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	t.Log(SearchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	t.Log(SearchInRotatedSortedArray([]int{1}, 0))
}
