package leetcode

import "testing"

// https://leetcode.cn/problems/minimum-path-sum/

func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)

	for i := range grid {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = MPathSum(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][columns-1]
}

func MPathSum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestMinPathSum(t *testing.T) {
	t.Log(MinPathSum([][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}))
	t.Log(MinPathSum([][]int{[]int{1, 2, 3}, []int{4, 5, 6}}))
}
