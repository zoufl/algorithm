package leetcode

import "testing"

// https://leetcode-cn.com/problems/set-matrix-zeroes/

// 记录0出现的位置，然后再遍历置位0
func SetZeroes(matrix [][]int) {
	r := make(map[int]bool)
	c := make(map[int]bool)

	for i, rows := range matrix {
		for j, item := range rows {
			if item == 0 {
				r[i] = true
				c[j] = true
			}
		}
	}

	for i, rows := range matrix {
		for j := range rows {
			if r[i] || c[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	SetZeroes(matrix)
	t.Log(matrix)

	matrix = [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	SetZeroes(matrix)
	t.Log(matrix)

	matrix = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 0, 7, 8},
		[]int{0, 10, 11, 12},
		[]int{13, 14, 15, 0},
	}
	SetZeroes(matrix)
	t.Log(matrix)
}
