package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	result := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}
