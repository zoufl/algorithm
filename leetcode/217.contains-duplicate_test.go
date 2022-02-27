package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/contains-duplicate/

// 使用哈希表标记
func ContainsDuplicate(nums []int) bool {
	hashTable := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := hashTable[num]; ok {
			return true
		}

		hashTable[num] = struct{}{}
	}

	return false
}

// 先排序，再找重复元素
func ContainsDuplicate1(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(ContainsDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	t.Log(ContainsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	t.Log(ContainsDuplicate(nums))
}
