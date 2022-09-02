package leetcode

import "testing"

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

// 二分查找 O(logN)
func SearchRange(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			i, j := mid, mid

			for i >= 0 && nums[i] == target {
				i--
			}

			for j <= n-1 && nums[j] == target {
				j++
			}

			return []int{i + 1, j - 1}
		}

		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return []int{-1, -1}
}

// 双指针
func SearchRange1(nums []int, target int) []int {
	n := len(nums)
	start, end := 0, n-1

	for start <= end {
		if target < nums[start] || target > nums[end] {
			return []int{-1, -1}
		} else if target == nums[start] && target == nums[end] {
			return []int{start, end}
		}

		if target > nums[start] {
			start++
		}

		if target < nums[end] {
			end--
		}

	}

	return []int{-1, -1}
}

func TestSearchRange(t *testing.T) {
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(SearchRange([]int{}, 0))
	t.Log(SearchRange([]int{1}, 1))
}
