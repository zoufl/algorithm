package leetcode

import "testing"

// https://leetcode.cn/problems/permutations/

func Permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	n := len(nums)

	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			dfs(index + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs(0)

	return ans
}

func TestPermute(t *testing.T) {
	t.Log(Permute([]int{1, 2, 3}))
	//t.Log(Permute([]int{0, 1}))
	//t.Log(Permute([]int{1}))
}
