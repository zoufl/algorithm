package leetcode

import (
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/valid-parentheses/

// 使用栈
func IsValidParentheses(s string) bool {
	parentheses := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			n := len(stack)
			if n == 0 || parentheses[stack[n-1]] != ch {
				return false
			}

			stack = stack[:n-1]
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
