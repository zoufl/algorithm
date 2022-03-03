package leetcode

import "testing"

// https://leetcode-cn.com/problems/rotate-image/

// 原地旋转
func RotateImage1(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 用翻转代替旋转，先水平翻转，然后主对角线翻转
func RotateImage(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
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
