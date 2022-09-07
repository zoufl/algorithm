package leetcode

import (
	"testing"
)

// https://leetcode.cn/problems/binary-tree-inorder-traversal/

func InorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	var inorder func(*TreeNode)
	inorder = func(t *TreeNode) {
		if t == nil {
			return
		}

		inorder(t.Left)
		ans = append(ans, t.Val)
		inorder(t.Right)
	}

	inorder(root)

	return ans
}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	t.Log(InorderTraversal(root))

	root = &TreeNode{}
	t.Log(InorderTraversal(root))

	root = &TreeNode{
		Val: 1,
	}
	t.Log(InorderTraversal(root))
}
