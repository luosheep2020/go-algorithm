package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
