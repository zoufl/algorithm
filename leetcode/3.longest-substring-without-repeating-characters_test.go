package leetcode

import "testing"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	n := len(s)
	ans := 0
	r := -1

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for r+1 < n && !m[s[r+1]] {
			m[s[r+1]] = true
			r++
		}

		ans = MaxLength(ans, len(m))
	}

	return ans
}

func MaxLength(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(LengthOfLongestSubstring("abcabcbb"))
	t.Log(LengthOfLongestSubstring("bbbbb"))
	t.Log(LengthOfLongestSubstring("pwwkew"))
	t.Log(LengthOfLongestSubstring(" "))
	t.Log(LengthOfLongestSubstring("au"))
	t.Log(LengthOfLongestSubstring("aab"))
}
