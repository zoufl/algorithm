package leetcode

import (
	"testing"
)

// https://leetcode.cn/problems/unique-paths/

// 动态规划
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
	t.Log(UniquePaths(3, 7))
	t.Log(UniquePaths(3, 2))
	t.Log(UniquePaths(7, 3))
	t.Log(UniquePaths(3, 3))
}
