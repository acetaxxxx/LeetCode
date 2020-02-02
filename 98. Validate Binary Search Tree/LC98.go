package lc98

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	// if node is null return true
	if root == nil {
		return true
	}
	// if node left not exist but right exist
	if root.Left == nil && root.Right != nil {
		return false
	}

	if root.Left != nil && root.Val < root.Left.Val {
		return false
	}

	if root.Right != nil && root.Val > root.Right.Val {
		return false
	}
	return isValidBST(root.Left) || isValidBST(root.Right)
}
