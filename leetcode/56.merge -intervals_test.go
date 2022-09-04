package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/merge-intervals/

// 先排序，双指针
func MergeIntervals(intervals [][]int) [][]int {
	n := len(intervals)

	if n == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		l, r := intervals[i][0], intervals[i][1]

		if len(ans) == 0 || ans[len(ans)-1][1] < l {
			ans = append(ans, []int{l, r})
		} else {
			ans[len(ans)-1][1] = MaxInterval(ans[len(ans)-1][1], r)
		}
	}

	return ans
}

func MaxInterval(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestMergeIntervals(t *testing.T) {
	t.Log(MergeIntervals([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	t.Log(MergeIntervals([][]int{[]int{1, 4}, []int{4, 5}}))
}
