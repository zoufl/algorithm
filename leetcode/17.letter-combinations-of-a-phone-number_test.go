package leetcode

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

var phoneMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

var combinations []string

// 回溯法
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations = make([]string, 0)
	backtrack(digits, "")
	return combinations
}

func backtrack(digits string, combination string) {
	if len(digits) == 0 {
		combinations = append(combinations, combination)
	} else {
		letters := phoneMap[string(digits[0])]
		for _, letter := range letters {
			backtrack(digits[1:], combination+letter)
		}
	}
}

// 队列法
func LetterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	queue := make([]string, 0)
	dArr := strings.Split(digits, "")

	for _, num := range dArr {
		letters := phoneMap[num]

		if len(queue) == 0 {
			queue = append(queue, letters...)
			continue
		}

		qn := len(queue)
		for j := 0; j < qn; j++ {
			k := queue[0]
			for _, letter := range letters {
				queue = append(queue, k+letter)
			}
			queue = queue[1:]
		}
	}

	return queue
}

// map记录
func LetterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := make(map[string]int)
	dArr := strings.Split(digits, "")
	level := 0

	for i, num := range dArr {
		letters := phoneMap[num]
		for _, letter := range letters {
			if i == 0 {
				m[letter] = i
			} else {
				copiedM := make(map[string]int)
				for k, v := range m {
					if v == i-1 {
						copiedM[k] = v
					}
				}

				for k := range copiedM {
					m[k+letter] = i

				}
			}

			level = i
		}
	}

	ans := make([]string, 0)
	for k, v := range m {
		if v == level {
			ans = append(ans, k)
		}
	}

	return ans
}

func TestLetterCombinations(t *testing.T) {
	t.Log(LetterCombinations("23"))
	//t.Log(LetterCombinations("234"))
	//t.Log(LetterCombinations(""))
	//t.Log(LetterCombinations("2"))
}
