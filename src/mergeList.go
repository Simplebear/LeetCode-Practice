package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(list1, list2 *ListNode) *ListNode {
	var result *ListNode = &ListNode{Val: 0}
	temp := result
	for list1.Next != nil && list2.Next != nil {
		if list1.Val < list2.Val {
			temp.Val = list1.Val
			list1 = list1.Next
		} else {
			temp.Val = list2.Val
			list1 = list2.Next
		}
		temp = temp.Next
	}
	return result
}
