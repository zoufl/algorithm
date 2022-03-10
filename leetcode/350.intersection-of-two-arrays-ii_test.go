package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

// 哈希表
func IntersectionII1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ans := make([]int, 0)

	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if m[num] > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}

	return ans
}

// 双指针
func IntersectionII(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	ans := make([]int, 0)
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}

func TestIntersectionII(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	t.Log(IntersectionII(nums1, nums2))

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	t.Log(IntersectionII(nums1, nums2))
}
