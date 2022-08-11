package tree

// TreeNode is defined for a binary tree node.
type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
