package leetcode

import "testing"

// https://leetcode-cn.com/problems/jewels-and-stones/

func NumJewelsInStones(jewels string, stones string) int {
	if len(jewels) == 0 || len(stones) == 0 {
		return 0
	}

	m := make(map[rune]struct{})
	ans := 0

	for _, j := range jewels {
		m[j] = struct{}{}
	}

	for _, s := range stones {
		if _, ok := m[s]; ok {
			ans++
		}
	}

	return ans
}

func TestNumJewelsInStones(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	t.Log(NumJewelsInStones(jewels, stones))

	jewels = "z"
	stones = "ZZ"
	t.Log(NumJewelsInStones(jewels, stones))
}
