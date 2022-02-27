package leetcode

import "testing"

// https://leetcode-cn.com/problems/plus-one/

// 整数进位情况
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1

			return digits
		}
	}

	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1

	return newDigits
}

func TestPlusOne(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(PlusOne(nums))

	nums = []int{4, 3, 2, 1}
	t.Log(PlusOne(nums))

	nums = []int{9, 9}
	t.Log(PlusOne(nums))

	nums = []int{0}
	t.Log(PlusOne(nums))
}
