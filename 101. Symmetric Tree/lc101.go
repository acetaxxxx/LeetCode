package lc101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isLevelSymmetric(root.Left, root.Right)
}

func isLevelSymmetric(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val == right.Val {
			return isLevelSymmetric(left.Left, right.Right) && isLevelSymmetric(left.Right, right.Left)
		}
	} else if left == nil && right == nil {
		return true
	}
	return false
}
