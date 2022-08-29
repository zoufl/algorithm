package leetcode

import "testing"

// https://leetcode.cn/problems/longest-palindromic-substring/

// 中心扩散法
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}

		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

// 暴力法
func LongestPalindrome1(s string) string {
	var ss string

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] {
				if isPalindromes(s[i:j+1]) && len(ss) < j-i+1 {
					ss = s[i : j+1]
				}
			}
		}
	}

	return ss
}

func isPalindromes(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func TestLongestPalindrome(t *testing.T) {
	t.Log(LongestPalindrome("babad"))
	t.Log(LongestPalindrome("cbbd"))
	t.Log(LongestPalindrome("a"))
}
