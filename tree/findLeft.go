package main

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	findLeft(root, &res)
	return res
}
func findLeft(root *TreeNode, res *int) {
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}
	if root.Left != nil {
		findLeft(root.Left, res)
	}
	if root.Right != nil {
		findLeft(root.Right, res)
	}
}
