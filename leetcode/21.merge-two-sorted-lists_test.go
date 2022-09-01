package leetcode

import "testing"

// https://leetcode.cn/problems/merge-two-sorted-lists/

// 迭代
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	prev := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	} else if list2 == nil {
		prev.Next = list1
	}

	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	t.Log(MergeTwoLists(l1, l2))

	t.Log(MergeTwoLists(nil, nil))

	t.Log(MergeTwoLists(nil, &ListNode{}))

	l1 = &ListNode{
		Val:  2,
		Next: nil,
	}

	l2 = &ListNode{
		Val:  1,
		Next: nil,
	}
	t.Log(MergeTwoLists(l1, l2))

	l1 = &ListNode{
		Val:  5,
		Next: nil,
	}

	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	t.Log(MergeTwoLists(l1, l2))
}
