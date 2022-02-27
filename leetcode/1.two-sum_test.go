package leetcode

import "testing"

// https://leetcode-cn.com/problems/two-sum/

// 哈希表记录元素和index
func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if another, ok := hashTable[target-num]; ok {
			return []int{another, i}
		}

		hashTable[num] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	t.Log(TwoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	t.Log(TwoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	t.Log(TwoSum(nums, target))

}
