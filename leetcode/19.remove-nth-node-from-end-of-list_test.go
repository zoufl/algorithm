package leetcode

import "testing"

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// 快慢双指针
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	first, second := head, newHead

	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}

	second.Next = second.Next.Next

	return newHead.Next
}

// 计算链表长度
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	length := 0
	h := head

	for h != nil {
		length++
		h = h.Next
	}

	h = head
	idx := length - n
	if idx == 0 {
		head = h.Next
	} else {
		for i := 0; i < length; i++ {
			if i == idx-1 {
				h.Next = h.Next.Next
				break
			}

			h = h.Next
		}
	}

	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	t.Log(RemoveNthFromEnd(l, 2))

	l = &ListNode{
		Val:  1,
		Next: nil,
	}
	t.Log(RemoveNthFromEnd(l, 1))

	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	t.Log(RemoveNthFromEnd(l, 1))

	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	t.Log(RemoveNthFromEnd(l, 2))
}
