package leetcode

import (
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/valid-parentheses/

// 使用栈
func IsValidParentheses(s string) bool {
	stack := make([]rune, 0)
	m := make(map[rune]rune)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'

	for _, c := range s {
		if '(' == c || '{' == c || '[' == c {
			stack = append(stack, c)
		} else if len(stack) > 0 && m[stack[len(stack)-1]] == c {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func IsValidParentheses1(s string) bool {
	for len(s) > 0 {
		n := len(s)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)

		if len(s) == n {
			break
		}
	}

	return len(s) == 0
}

func TestIsValidParentheses(t *testing.T) {
	s := "()"
	t.Log(IsValidParentheses(s))

	s = "()[]{}"
	t.Log(IsValidParentheses(s))

	s = "(]"
	t.Log(IsValidParentheses(s))

	s = "([)]"
	t.Log(IsValidParentheses(s))

	s = "{[]}"
	t.Log(IsValidParentheses(s))
}
