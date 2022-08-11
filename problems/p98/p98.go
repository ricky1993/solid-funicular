package p98

import "github.com/ricky1993/solid-funicular/structures/tree"

func max(root *tree.TreeNode[int]) int {
	if root.Right == nil {
		return root.Val
	}
	return max(root.Right)
}

func min(root *tree.TreeNode[int]) int {
	if root.Left == nil {
		return root.Val
	}
	return min(root.Left)
}

func isValidBST(root *tree.TreeNode[int]) bool {
	if root == nil {
		return true
	}
	right := false
	left := false
	if root.Right == nil {
		right = true
	} else {
		suc := min(root.Right)
		if root.Val < suc {
			right = isValidBST(root.Right)
		}
	}
	if root.Left == nil {
		left = true
	} else {
		pre := max(root.Left)
		if pre < root.Val {
			left = isValidBST(root.Left)
		}
	}
	return left && right
}
