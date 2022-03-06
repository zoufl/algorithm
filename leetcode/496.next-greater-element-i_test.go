package leetcode

import "testing"

// https://leetcode-cn.com/problems/next-greater-element-i/

func NextGreaterElement1(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	m := make(map[int]int)

	for i, num2 := range nums2 {
		if _, ok := m[num2]; !ok {
			m[num2] = i
		}
	}

	for i, num1 := range nums1 {
		index := m[num1]
		j := index + 1

		for ; j < len(nums2); j++ {
			if nums2[j] > num1 {
				ans[i] = nums2[j]
				break

			}
		}

		if j == len(nums2) {
			ans[i] = -1
		}
	}

	return ans
}

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	m := make(map[int]int)
	var stack []int

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if v, ok := m[num]; ok {
			ans[i] = v
		} else {
			ans[i] = -1
		}
	}

	return ans
}

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	t.Log(NextGreaterElement(nums1, nums2))

	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	t.Log(NextGreaterElement(nums1, nums2))
}
