package leetcode

import (
	"sort"
	"strings"
	"testing"
)

func IsValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		if _, ok := m[c]; !ok {
			return false
		}

		m[c]--

		if m[c] == 0 {
			delete(m, c)
		}
	}

	return len(m) == 0
}

// 直接排序后比较
func IsValidAnagram1(s string, t string) bool {
	return sortStr(s) == sortStr(t)
}

func sortStr(s string) string {
	r := strings.Split(s, "")
	sort.Strings(r)
	return strings.Join(r, "")
}

func TestIsValidAnagram(t *testing.T) {
	s := "anagram"
	t1 := "nagaram"
	t.Log(IsValidAnagram(s, t1))

	s = "rat"
	t1 = "car"
	t.Log(IsValidAnagram(s, t1))
}
