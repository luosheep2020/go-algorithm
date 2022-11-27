package main

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}
