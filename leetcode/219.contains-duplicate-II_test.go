package leetcode

import (
	"math"
	"testing"
)

// https://leetcode-cn.com/problems/contains-duplicate-ii/

func containsNearbyDuplicate(nums []int, k int) bool {
	hashTable := map[int]int{}
	for i, num := range nums {
		if j, ok := hashTable[num]; ok && math.Abs(float64(j-i)) <= float64(k) {
			return true
		}

		hashTable[num] = i
	}

	return false
}

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	t.Log(containsNearbyDuplicate(nums, k))

	nums = []int{1, 0, 1, 1}
	k = 1
	t.Log(containsNearbyDuplicate(nums, k))

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	t.Log(containsNearbyDuplicate(nums, k))
}
