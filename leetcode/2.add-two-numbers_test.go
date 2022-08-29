package leetcode

import "testing"

// https://leetcode.cn/problems/add-two-numbers/

// 双指针
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	add := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{}
		}

		if l2 == nil {
			l2 = &ListNode{}
		}

		val := l1.Val + l2.Val + add
		add = val / 10
		val %= 10

		node := &ListNode{
			Val: val,
		}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if add > 0 {
		tail.Next = &ListNode{Val: add}
	}

	return head
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	t.Log(AddTwoNumbers(l1, l2))

	l1 = &ListNode{
		Val:  0,
		Next: nil,
	}

	l2 = &ListNode{
		Val:  0,
		Next: nil,
	}

	t.Log(AddTwoNumbers(l1, l2))

	l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	l2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	t.Log(AddTwoNumbers(l1, l2))
}
