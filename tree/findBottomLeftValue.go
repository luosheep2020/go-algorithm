package main

var maxDeep int // 全局变量 深度
var value int   //全局变量 最终值

func findLeftValue(root *TreeNode, deep int) {
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			value = root.Val
			maxDeep = deep
		}
	}

	if root.Left != nil {
		deep++
		findLeftValue(root.Left, deep)
		deep--
	}
	if root.Right != nil {
		deep++
		findLeftValue(root.Right, deep)
		deep--
	}
}
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	findLeftValue(root, maxDeep)
	return value
}
