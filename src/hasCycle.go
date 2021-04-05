package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(l *ListNode) bool {
	if l == nil || l.Next == nil {
		return false
	}
	slow := l
	fast := l.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
