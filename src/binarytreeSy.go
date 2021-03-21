package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(n, m *TreeNode) bool {
	if n == nil && m == nil {
		return true
	}
	if n == nil || m == nil {
		return false
	}
	return n.Val == m.Val && check(n.Left, m.Right) && check(n.Right, m.Left)
}
