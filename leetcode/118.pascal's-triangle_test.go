package leetcode

import "testing"

// https://leetcode-cn.com/problems/pascals-triangle/submissions/

func PascalTriangle(numRows int) [][]int {
	rows := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		rows[i] = make([]int, i+1)

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				rows[i][j] = 1

				continue
			}

			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
	}

	return rows
}

func TestPascalTriangle(t *testing.T) {
	row := 5
	t.Log(PascalTriangle(row))

	row = 9
	t.Log(PascalTriangle(row))
}
