package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/group-anagrams/

// 排序
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)

		//r := strings.Split(str, "")
		//sort.Strings(r)
		//s := strings.Join(r, "")

		if _, ok := m[sortedStr]; ok {
			m[sortedStr] = append(m[sortedStr], str)
		} else {
			m[sortedStr] = []string{str}
		}
	}

	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(GroupAnagrams(strs))

	strs = []string{""}
	t.Log(GroupAnagrams(strs))

	strs = []string{"a"}
	t.Log(GroupAnagrams(strs))
}
