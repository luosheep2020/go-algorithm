package main

func findMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	index := findMax(nums)

	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return root
}
