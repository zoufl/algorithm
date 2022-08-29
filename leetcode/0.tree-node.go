package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func (t *TreeNode) String() string {
//	var sBuilder strings.Builder
//
//	var curr = l
//	for curr != nil {
//		sBuilder.WriteString(strconv.Itoa(curr.Val))
//		if curr.Next != nil {
//			sBuilder.WriteString("->")
//		}
//		curr = curr.Next
//	}
//
//	return sBuilder.String()
//}
