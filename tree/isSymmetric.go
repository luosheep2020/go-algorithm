package main

func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}

func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}
