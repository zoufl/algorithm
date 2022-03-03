package leetcode

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/valid-sudoku/

func IsValidSudoku(board [][]byte) bool {
	m3 := make(map[string]struct{})

	for i := 0; i < len(board); i++ {
		m1 := make(map[string]struct{})
		m2 := make(map[string]struct{})

		for j := 0; j < len(board[i]); j++ {
			k1 := string(board[i][j])
			if k1 != "." {
				if _, ok := m1[k1]; ok {
					return false
				}

				m1[k1] = struct{}{}
			}

			k2 := string(board[j][i])
			if k2 != "." {
				if _, ok := m2[k2]; ok {
					return false
				}

				m2[k2] = struct{}{}
			}

			k3 := fmt.Sprintf("%d_%d_%d", i/3, j/3, board[i][j])
			if board[i][j] != '.' {
				if _, ok := m3[k3]; ok {
					return false
				}

				m3[k3] = struct{}{}
			}
		}
	}

	return true
}

func IsValidSudoku1(board [][]byte) bool {
	var rows, columns [9][9]int
	var subBoxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subBoxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	t.Log(IsValidSudoku(board))

	board = [][]byte{
		[]byte("83..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("1...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	t.Log(IsValidSudoku(board))
}
