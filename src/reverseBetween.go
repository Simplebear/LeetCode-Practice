package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: 0}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curl := pre.Next
	for i := 0; i < right-left; i++ {
		next := curl.Next
		curl.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode
}
