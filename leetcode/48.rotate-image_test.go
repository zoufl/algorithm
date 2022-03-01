package leetcode

import "testing"

// https://leetcode-cn.com/problems/rotate-image/

func RotateImage(matrix [][]int) {

}

func TestRotateImage(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	RotateImage(matrix)
	t.Log(matrix)

	matrix = [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	RotateImage(matrix)
	t.Log(matrix)
}
