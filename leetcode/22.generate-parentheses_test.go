package leetcode

import "testing"

//https://leetcode.cn/problems/generate-parentheses/

// 回溯法
func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)

	if n <= 0 {
		return ans
	}

	var dfs func(string, int, int)
	dfs = func(path string, open, close int) {
		if len(path) == 2*n {
			ans = append(ans, path)
			return
		}

		if open < n {
			dfs(path+"(", open+1, close)
		}

		if close < open {
			dfs(path+")", open, close+1)
		}
	}

	dfs("", 0, 0)
	return ans
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(GenerateParenthesis(3))
	t.Log(GenerateParenthesis(1))
}
