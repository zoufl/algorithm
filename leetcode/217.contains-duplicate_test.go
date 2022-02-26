package leetcode

import "testing"

// https://leetcode-cn.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	hashTable := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := hashTable[num]; ok {
			return true
		}

		hashTable[num] = struct{}{}
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(containsDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	t.Log(containsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	t.Log(containsDuplicate(nums))
}
