package leetcode

import "testing"

// https://leetcode.cn/problems/climbing-stairs/

// 动态规划
func ClimbStairs(n int) int {
	p, q, r := 0, 0, 1

	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}

func ClimbStairs1(n int) int {
	ans := 0

	var dfs func(int)
	dfs = func(target int) {
		if target < 0 {
			return
		}

		if target == 0 {
			ans++
			return
		}

		for i := 1; i <= 2; i++ {
			dfs(target - i)
		}
	}

	dfs(n)

	return ans
}

func TestClimbStairs(t *testing.T) {
	t.Log(ClimbStairs(2))
	t.Log(ClimbStairs(3))
	t.Log(ClimbStairs(4))
	t.Log(ClimbStairs(5))
	t.Log(ClimbStairs(44))
}
