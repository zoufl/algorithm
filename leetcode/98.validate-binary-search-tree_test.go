package leetcode

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/validate-binary-search-tree/

// 先序搜索
func IsValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}

		if node.Val <= min || node.Val >= max {
			return false
		}

		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
}

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t.Log(IsValidBST(root))

	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	t.Log(IsValidBST(root))

	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	t.Log(IsValidBST(root))

	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	t.Log(IsValidBST(root))
}
