package leetcode

import (
	"sort"
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

// 滑动窗口+双指针
func FindAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	ans := make([]int, 0)

	if sLen < pLen {
		return ans
	}

	var sCount, pCount [26]int
	for _, c := range p {
		pCount[c-'a']++
	}

	for l, r := 0, 0; r < sLen; r++ {
		sCount[s[r]-'a']++
		if r-l+1 > pLen {
			sCount[s[l]-'a']--
			l++
		}

		if sCount == pCount {
			ans = append(ans, l)
		}
	}

	return ans
}

// 滑动窗口
func FindAnagrams2(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	ans := make([]int, 0)

	if sLen < pLen {
		return ans
	}

	var sCount, pCount [26]int
	for i, c := range p {
		sCount[s[i]-'a']++
		pCount[c-'a']++
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	//for i, c := range s[:sLen-pLen] {
	//	sCount[c-'a']--
	//	sCount[s[i+pLen]-'a']++
	//	if sCount == pCount {
	//		ans = append(ans, i+1)
	//	}
	//}

	for i := 0; i < sLen-pLen; i++ {
		sCount[s[i]-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return ans

}

// 执行超时
func FindAnagrams1(s string, p string) []int {
	k := len(p)
	r := strings.Split(p, "")
	sort.Strings(r)
	sortedStr := strings.Join(r, "")
	var ans []int

	for i := 0; i <= len(s)-k; i++ {
		tmp := s[i : i+k]
		ss := strings.Split(tmp, "")
		sort.Strings(ss)
		sorted := strings.Join(ss, "")

		if sorted == sortedStr {
			ans = append(ans, i)
		}
	}

	return ans
}

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	t.Log(FindAnagrams(s, p))

	s = "abab"
	p = "ab"
	t.Log(FindAnagrams(s, p))
}
