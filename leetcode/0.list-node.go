package leetcode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var sBuilder strings.Builder

	var curr = l
	for curr != nil {
		sBuilder.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			sBuilder.WriteString("->")
		}
		curr = curr.Next
	}

	return sBuilder.String()
}
