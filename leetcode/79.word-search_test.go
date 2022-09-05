package leetcode

import "testing"

// https://leetcode.cn/problems/word-search/

func ExistWord(board [][]byte, word string) bool {
	return true
}

func TestExistWord(t *testing.T) {
	t.Log(ExistWord([][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCCED"))
	t.Log(ExistWord([][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "SEE"))
	t.Log(ExistWord([][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, "ABCB"))
}
