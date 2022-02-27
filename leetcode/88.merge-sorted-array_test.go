package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/merge-sorted-array/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0

	for i, num := range nums1 {
		if num == 0 {
			nums1[i] = nums2[j]
			j++
		}

		if j >= n {
			break
		}
	}

	sort.Ints(nums1)
}

// 复制，排序
func Merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:])
	sort.Ints(nums1)
}

// 倒序比较，双指针
func Merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	Merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	Merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	Merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m = 6
	nums2 = []int{1, 2, 2}
	n = 3
	Merge(nums1, m, nums2, n)
	t.Log(nums1)
}
