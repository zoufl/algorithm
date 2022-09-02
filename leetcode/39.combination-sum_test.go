package leetcode

import "testing"

// https://leetcode.cn/problems/combination-sum/

// 回溯法
func CombinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int)

	dfs = func(target, index int) {
		if target < 0 {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)

	return ans
}

// 回溯法
func CombinationSum1(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var backtrack func([]int, int)
	backtrack = func(combinations []int, target int) {
		sum := 0
		for _, combination := range combinations {
			sum += combination
		}

		if sum == target {
			ans = append(ans, combinations)
			return
		}

		if sum < target {
			for _, candidate := range candidates {
				if !gtAll(combinations, candidate) {
					continue
				}

				res := sum + candidate

				if res <= target {
					var newCombinations []int
					newCombinations = append(newCombinations, append([]int{}, combinations...)...)
					newCombinations = append(newCombinations, candidate)
					backtrack(newCombinations, target)
				}
			}
		}
	}

	backtrack([]int{}, target)

	return ans
}

func gtAll(nums []int, target int) bool {
	for _, n := range nums {
		if target < n {
			return false
		}
	}

	return true
}

func TestCombinationSum(t *testing.T) {
	t.Log(CombinationSum([]int{2, 3, 6, 7}, 7))
	t.Log(CombinationSum([]int{2, 3, 5}, 8))
	t.Log(CombinationSum([]int{2}, 1))
	t.Log(CombinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}
