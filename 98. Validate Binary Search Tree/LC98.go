package lc98

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "math"

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

	return validate(root, math.Inf(-1), math.Inf(1))
}

func validate(root *TreeNode, minValue float64, maxValue float64) bool {
	if root == nil {
		return true
	}
	if float64(root.Val) <= minValue || float64(root.Val) >= maxValue {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return validate(root.Left, minValue,float64( root.Val)) && validate(root.Right, float64(root.Val), maxValue)
}
