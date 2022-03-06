package leetcode

import "testing"

// https://leetcode-cn.com/problems/score-of-parentheses/

func ScoreOfParentheses(s string) int {
	var stack []interface{}

	for _, val := range s {
		if val == '(' {
			stack = append(stack, val)
		} else {
			k := 0
			for stack[len(stack)-1] != '(' {
				k += stack[len(stack)-1].(int)
				stack = stack[:len(stack)-1]
			}

			if k == 0 {
				stack[len(stack)-1] = 1
			} else {
				stack[len(stack)-1] = 2 * k
			}
		}
	}

	res := 0
	for _, v := range stack {
		res += v.(int)
	}

	return res
}

func ScoreOfParentheses1(s string) int {
	ans, bal := 0, 0

	for i, c := range s {
		if c == '(' {
			bal++
		} else {
			bal--

			if s[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}

	return ans
}

func TestScoreOfParentheses(t *testing.T) {
	s := "()"
	t.Log(ScoreOfParentheses(s))

	s = "(())"
	t.Log(ScoreOfParentheses(s))

	s = "()()"
	t.Log(ScoreOfParentheses(s))

	s = "(()(()))"
	t.Log(ScoreOfParentheses(s))

	s = "((((((())))()())))"
	t.Log(ScoreOfParentheses(s))
}
