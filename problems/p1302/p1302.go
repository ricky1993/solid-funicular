package p1302

import "github.com/ricky1993/solid-funicular/structures/tree"

func deepestLeavesSum(root *tree.TreeNode[int]) int {
	deepestSum := 0
	deepest := 0
	deepestSumOfLeaves(root, &deepest, &deepestSum, 0)
	return deepestSum
}

func deepestSumOfLeaves(root *tree.TreeNode[int], deepest, deepestSum *int, depth int) {
	if root == nil {
		return
	}
	if depth > *deepest {
		*deepest = depth
		*deepestSum = root.Val
	} else if depth == *deepest {
		*deepestSum += root.Val
	}
	deepestSumOfLeaves(root.Left, deepest, deepestSum, depth+1)
	deepestSumOfLeaves(root.Right, deepest, deepestSum, depth+1)
}
