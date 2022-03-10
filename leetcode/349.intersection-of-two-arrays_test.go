package leetcode

import "testing"

// https://leetcode-cn.com/problems/intersection-of-two-arrays/

// 哈希表
func Intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	ans := make([]int, 0)

	for _, num := range nums1 {
		m[num] = true
	}

	for _, num := range nums2 {
		if m[num] {
			m[num] = false
			ans = append(ans, num)
		}
	}

	return ans
}

func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	t.Log(Intersection(nums1, nums2))

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	t.Log(Intersection(nums1, nums2))
}
