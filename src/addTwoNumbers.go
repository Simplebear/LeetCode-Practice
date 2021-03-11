package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//var res *ListNode = &ListNode{Val: 0}
	var r *ListNode = &ListNode{Val: 0}
	res := r
	if l1 == nil && l2 != nil {
		res.Val = l2.Val
		res.Next = l2.Next
	}
	if l1 != nil && l2 == nil {
		res.Val = l1.Val
		res.Next = l1.Next
	}
	if l1 != nil && l2 != nil {
		res.Val = l1.Val + l2.Val
		res.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	for res != nil {
		c := res.Val - 10
		if c >= 0 {
			if res.Next != nil {
				res.Next.Val += 1
				res.Val = c
			} else {
				res.Val = c
				res.Next = &ListNode{Val: 1}
			}
		}
		if res.Next != nil && res.Next.Val == 0 && res.Next.Next == nil {
			res.Next = nil
		}
		res = res.Next

	}

	return r
}
