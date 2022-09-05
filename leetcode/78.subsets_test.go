package leetcode

import "testing"

// https://leetcode.cn/problems/subsets/

// 回溯
func Subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(index int) {
		ans = append(ans, append([]int{}, path...))

		for i := index; i < n; i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return ans
}

func TestSubsets(t *testing.T) {
	t.Log(Subsets([]int{1, 2, 3}))
	t.Log(Subsets([]int{0}))
}
