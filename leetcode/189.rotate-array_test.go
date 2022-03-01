package leetcode

import "testing"

// https://leetcode-cn.com/problems/rotate-array/

func RotateArray1(nums []int, k int) {
	newArray := make([]int, len(nums))

	for i, num := range nums {
		newArray[(i+k)%len(newArray)] = num
	}

	copy(nums, newArray)
}

// 数组翻转，整体翻转，左部分翻转，右部分翻转
func RotateArray(nums []int, k int) {
	k %= len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	RotateArray(nums, k)
	t.Log(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	RotateArray(nums, k)
	t.Log(nums)
}
