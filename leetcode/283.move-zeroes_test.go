package leetcode

import "testing"

// https://leetcode-cn.com/problems/move-zeroes/

//func moreZeroes1(nums []int) {
//	j := 0
//
//	for i := 0; i < len(nums); i++ {
//		if (nums)[i] != 0 {
//			continue
//		}
//
//		if j < i {
//			j = i + 1
//		}
//
//		for ; j < len(nums); j++ {
//			if (nums)[j] != 0 {
//				(nums)[i], (nums)[j] = (nums)[j], (nums)[i]
//				break
//			}
//		}
//	}
//}

func moreZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if (nums)[j] != 0 {
			if i != j {
				(nums)[i], (nums)[j] = (nums)[j], (nums)[i]
			}
			i++
		}
	}
}

func TestMoreZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moreZeroes(nums)
	t.Log(nums)

	nums = []int{0}
	moreZeroes(nums)
	t.Log(nums)
}
